// @Title
// @Description
// @Author  Wangwengang  2023/12/22 13:11
// @Update  Wangwengang  2023/12/22 13:11
package core

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zutils"
	"github.com/smallnest/rpcx/protocol"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/proto/logic"
	"os"
)

type GlobalManager struct {
	nt *Notify

	ServerId int64
	IDWorker *zutils.IDWorker
}

var GlobalMgr *GlobalManager

func NewGlobalMgr(serverId int64) {
	IDWorker, err := zutils.NewIDWorker(serverId)
	if err != nil {
		os.Exit(1)
	}
	GlobalMgr = &GlobalManager{
		nt:       NewNotify(),
		ServerId: serverId,
		IDWorker: IDWorker,
	}
}

// 广播消息 发给所有用户
func (g *GlobalManager) SendMsgToAll(dataJson []byte, dataProto []byte) {
	_ = g.nt.NotifyBuffAllAuto(dataProto, dataJson)
}

// 消息发给指定连接
func (g *GlobalManager) SendMsgToConn(cID uint64, dataJson []byte, dataProto []byte) {
	if err := g.nt.NotifyBuffToConnByIDAuto(cID, dataJson, dataProto); err != nil {
		global.LOG.Errorf("NotifyBuffToConnByIDAuto err = %s", err.Error())
	}
}

func (g *GlobalManager) ConnectionEnterFirst(conn ziface.IConnection) {
	nextId, _ := g.IDWorker.NextID() // 雪花id
	conn.SetProperty("uniqueId", nextId)

	if g.nt.HasIdConn(uint64(nextId)) {
		g.nt.SetNotifyID(uint64(nextId), conn)
	}

	msg := logic.GateTologicMsg{
		From:     nextId,
		ServerId: GlobalMgr.ServerId,
	}
	data, _ := msg.Marshal()
	_, _, err := global.SRPC.RPC("Logic", "ConnectionBegin", data, protocol.ProtoBuffer, true)
	if err != nil {
		//如果logic服务不可用，将返回err
		global.LOG.Errorf("ConnectionEnterFirst Logic ConnectionBegin RPC err = %s", err.Error())
		conn.Stop()
	}

}

func (g *GlobalManager) ConnectionLost(conn ziface.IConnection) {
	if v, err := conn.GetProperty("uniqueId"); err == nil {
		g.nt.DelNotifyByID(uint64(v.(int64)))
		msg := logic.GateTologicMsg{
			From:     v.(int64),
			ServerId: GlobalMgr.ServerId,
		}
		data, _ := msg.Marshal()
		_, _, err := global.SRPC.RPC("Logic", "ConnectionLost", data, protocol.ProtoBuffer, true)
		if err != nil {
			//如果logic服务不可用，将返回err
			global.LOG.Errorf("JsonHandle RPC err = %s", err.Error())
		}
	}

	// hash桶中删除数据
}
