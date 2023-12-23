// @Title
// @Description
// @Author  Wangwengang  2023/12/22 20:55
// @Update  Wangwengang  2023/12/22 20:55
package core

import (
	"errors"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/zutils"
	"github.com/wwengg/im/proto/pbgate"
	"strconv"
)

type Notify struct {
	connIdMap zutils.ShardLockMaps
}

func NewNotify() *Notify {
	return &Notify{
		connIdMap: zutils.NewShardLockMaps(),
	}
}

func (n *Notify) genConnStrId(connID uint64) string {
	strConnId := strconv.FormatUint(connID, 10)
	return strConnId
}

func (n *Notify) ConnNums() int {
	return n.connIdMap.Count()
}

func (n *Notify) HasIdConn(Id uint64) bool {
	strId := n.genConnStrId(Id)
	return n.connIdMap.Has(strId)
}

func (n *Notify) SetNotifyID(Id uint64, conn ziface.IConnection) {
	strId := n.genConnStrId(Id)
	n.connIdMap.Set(strId, conn)
}

func (n *Notify) GetNotifyByID(Id uint64) (ziface.IConnection, error) {

	strId := n.genConnStrId(Id)
	Conn, ok := n.connIdMap.Get(strId)
	if !ok {
		return nil, errors.New(" Not Find UserId")
	}
	return Conn.(ziface.IConnection), nil
}

func (n *Notify) DelNotifyByID(Id uint64) {
	strId := n.genConnStrId(Id)
	n.connIdMap.Remove(strId)
}

func (n *Notify) NotifyToConnByID(Id uint64, MsgId uint32, data []byte) error {
	Conn, err := n.GetNotifyByID(Id)
	if err != nil {
		return err
	}
	err = Conn.SendMsg(MsgId, data)
	if err != nil {
		fmt.Printf("Notify to %d err:%s \n", Id, err)
		return err
	}
	return nil
}

func (n *Notify) NotifyAll(MsgId uint32, data []byte) error {

	n.connIdMap.IterCb(func(key string, v interface{}) {
		conn, _ := v.(ziface.IConnection)
		err := conn.SendMsg(MsgId, data)
		if err != nil {
			zlog.Ins().ErrorF("Notify to %s err:%s \n", key, err)
		}
	})

	return nil
}

func (n *Notify) NotifyBuffToConnByID(Id uint64, MsgId uint32, data []byte) error {
	Conn, err := n.GetNotifyByID(Id)
	if err != nil {
		return err
	}
	err = Conn.SendBuffMsg(MsgId, data)
	if err != nil {
		zlog.Ins().ErrorF("Notify to %d err:%s \n", Id, err)
		return err
	}
	return nil
}

func (n *Notify) NotifyBuffAll(MsgId uint32, data []byte) error {

	n.connIdMap.IterCb(func(key string, v interface{}) {
		conn, _ := v.(ziface.IConnection)
		err := conn.SendBuffMsg(MsgId, data)
		if err != nil {
			zlog.Ins().ErrorF("Notify to %s err:%s \n", key, err)
		}
	})

	return nil
}

func (n *Notify) NotifyBuffToConnByIDAuto(Id uint64, dataJson []byte, dataProto []byte) error {
	Conn, err := n.GetNotifyByID(Id)
	if err != nil {
		return err
	}
	if v, err := Conn.GetProperty("serializeType"); err == nil {
		serializeType, _ := v.(pbgate.Cmd)
		var data []byte
		switch serializeType {
		case pbgate.Cmd_JSON:
			data = dataJson
		case pbgate.Cmd_ProtoBuffer:
			data = dataProto
		default:
			data = dataProto
		}
		err := Conn.SendBuffMsg(uint32(serializeType), data)
		if err != nil {
			zlog.Ins().ErrorF("Notify to %d err:%s \n", Id, err)
		}
	}
	return nil
}

func (n *Notify) NotifyBuffAllAuto(dataProto []byte, dataJson []byte) error {

	n.connIdMap.IterCb(func(key string, v interface{}) {
		conn, _ := v.(ziface.IConnection)
		if v, err := conn.GetProperty("serializeType"); err == nil {
			serializeType, _ := v.(pbgate.Cmd)
			var data []byte
			switch serializeType {
			case pbgate.Cmd_JSON:
				data = dataJson
			case pbgate.Cmd_ProtoBuffer:
				data = dataProto
			default:
				data = dataProto
			}
			err := conn.SendBuffMsg(uint32(serializeType), data)
			if err != nil {
				zlog.Ins().ErrorF("Notify to %s err:%s \n", key, err)
			}
		}

	})

	return nil
}
