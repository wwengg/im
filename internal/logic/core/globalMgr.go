// @Title
// @Description
// @Author  Wangwengang  2024/1/8 12:40
// @Update  Wangwengang  2024/1/8 12:40
package core

import (
	"github.com/aceld/zinx/zutils"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/proto/logic"
	"github.com/wwengg/simple/core/store"
)

const REDIS_CONN_KEY = "conn_%d"

var GlobalMgr *GlobalManager

type GlobalManager struct {
	ConnMgr *ConnManager
	RoomMgr *RoomManager
	UserMgr *UserManager
}

func InitGlobalMgr() {
	GlobalMgr = &GlobalManager{
		RoomMgr: new(RoomManager),
		UserMgr: new(UserManager),
		ConnMgr: &ConnManager{
			connIdMap: zutils.ShardLockMaps{},
			RedisHash: store.NewRedisHash(global.RedisBase, REDIS_CONN_KEY),
		},
	}
}

func (g *GlobalManager) NewConn(args logic.GateTologicMsg) error {
	return g.ConnMgr.UpdateConnInfo(args)
}

func (g *GlobalManager) GetConn(args logic.GateTologicMsg) (Conn, error) {
	return g.ConnMgr.GetConnInfo(args)
}

func (g *GlobalManager) RemoveConn(args logic.GateTologicMsg) error {
	return g.ConnMgr.RemoveConnInfo(args)
}

func (g *GlobalManager) SendMsg() {

}
