// @Title
// @Description
// @Author  Wangwengang  2024/1/8 16:34
// @Update  Wangwengang  2024/1/8 16:34
package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aceld/zinx/zutils"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/proto/logic"
	"github.com/wwengg/simple/core/mutex"
	"github.com/wwengg/simple/core/store"
	"go.uber.org/zap"
)

const connMutexKey = "%s/logicConn/%d"

type Conn struct {
	CId       int64                             `json:"cId"`
	ServerId  int64                             `json:"serverId"`
	Serialize logic.GateTologicMsgSerializeType `json:"serialize"`
	UserId    int64                             `json:"userId,omitempty"`
	RoomId    int64                             `json:"roomId,omitempty"`
	AppId     int64                             `json:"appId,omitempty"`
}

type ConnManager struct {
	connIdMap zutils.ShardLockMaps
	RedisHash *store.RedisHash
}

// 更新连接数据到redis中
func (c *ConnManager) UpdateConnInfo(args logic.GateTologicMsg) error {
	global.LOG.Info("UpdateConnInfo Enter", zap.Any("args", args))
	if args.From == 0 && args.ServerId == 0 {
		global.LOG.Errorf("NewConn err", zap.Any("args", args))
		return errors.New("NewConn error")
	}
	mutexKey := fmt.Sprintf(connMutexKey, global.CONFIG.RPC.BasePath, args.From)
	if etcdMutex, err := mutex.NewEtcdMutex(mutexKey, global.ClientV3); err == nil {
		if err := etcdMutex.Lock(); err == nil {
			defer etcdMutex.Unlock()
		} else {
			global.LOG.Errorf("etcdMutex.Lock, err: %s", err.Error())
			return errors.New("etcdMutex.Lock error")
		}
	} else {
		global.LOG.Errorf("NewEtcdMutex err: %s", err.Error())
		return errors.New("NewEtcdMutex error")
	}

	info := Conn{
		CId:       args.From,
		ServerId:  args.ServerId,
		Serialize: args.Serialize,
	}
	global.LOG.Info("info", zap.Any("info", info))
	infoData, _ := json.Marshal(info)
	if err := c.RedisHash.HSet(info.CId, infoData); err != nil {

		return err
	}
	return nil
}

// 获取redis中的链接数据，如Serialize还是None，更新Serialize
func (c *ConnManager) GetConnInfo(args logic.GateTologicMsg) (Conn, error) {
	if args.From == 0 && args.ServerId == 0 {
		global.LOG.Errorf("NewConn err", zap.Any("args", args))
		return Conn{}, errors.New("NewConn error")
	}
	mutexKey := fmt.Sprintf(connMutexKey, global.CONFIG.RPC.BasePath, args.From)
	if etcdMutex, err := mutex.NewEtcdMutex(mutexKey, global.ClientV3); err == nil {
		if err := etcdMutex.Lock(); err == nil {
			defer etcdMutex.Unlock()
		} else {
			global.LOG.Errorf("etcdMutex.Lock, err: %s", err.Error())
			return Conn{}, errors.New("etcdMutex.Lock error")
		}
	} else {
		global.LOG.Errorf("NewEtcdMutex err: %s", err.Error())
		return Conn{}, errors.New("NewEtcdMutex error")
	}
	dataStr, err := c.RedisHash.HGet(args.From)
	if err != nil {
		global.LOG.Errorf("RedisHash.HGet args.From = %d,err=%s", args.From, err.Error())
		return Conn{}, err
	}
	var info Conn
	json.Unmarshal([]byte(dataStr), &info)
	if args.Serialize != logic.GateTologicMsg_None && info.Serialize == logic.GateTologicMsg_None {
		info.Serialize = args.Serialize
		infoData, _ := json.Marshal(info)
		if err := c.RedisHash.HSet(info.CId, infoData); err != nil {
			global.LOG.Errorf("RedisHash.HSet info = %v,err=%s", info, err.Error())
			return Conn{}, err
		}
	}
	return info, nil
}

func (c *ConnManager) RemoveConnInfo(args logic.GateTologicMsg) error {
	if args.From == 0 && args.ServerId == 0 {
		global.LOG.Errorf("NewConn err", zap.Any("args", args))
		return errors.New("NewConn error")
	}
	mutexKey := fmt.Sprintf(connMutexKey, global.CONFIG.RPC.BasePath, args.From)
	if etcdMutex, err := mutex.NewEtcdMutex(mutexKey, global.ClientV3); err == nil {
		if err := etcdMutex.Lock(); err == nil {
			defer etcdMutex.Unlock()
		} else {
			global.LOG.Errorf("etcdMutex.Lock, err: %s", err.Error())
			return errors.New("etcdMutex.Lock error")
		}
	} else {
		global.LOG.Errorf("NewEtcdMutex err: %s", err.Error())
		return errors.New("NewEtcdMutex error")
	}

	return c.RedisHash.HDel(args.From)
}
