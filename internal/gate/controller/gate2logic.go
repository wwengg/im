// @Title
// @Description
// @Author  Wangwengang  2023/12/22 12:52
// @Update  Wangwengang  2023/12/22 12:52
package controller

import (
	"github.com/aceld/zinx/ziface"
	"github.com/smallnest/rpcx/protocol"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/gate/core"
	"github.com/wwengg/im/proto/logic"
)

const (
	servivePath   = "Logic"
	serviceMethod = "GateToLogic"
)

func JsonHandle(request ziface.IRequest) {
	// oneWay true  shot
	if v, err := request.GetConnection().GetProperty("uniqueId"); err == nil {
		gateTologicMsg := logic.GateTologicMsg{
			From:      v.(int64),
			ServerId:  core.GlobalMgr.ServerId,
			Serialize: logic.GateTologicMsg_Json,
			Data:      request.GetData(),
		}

		data, _ := gateTologicMsg.Marshal()

		_, _, err := global.SRPC.RPC(servivePath, serviceMethod, data, protocol.ProtoBuffer, true)
		if err != nil {
			//如果logic服务不可用，将返回err
			global.LOG.Errorf("JsonHandle RPC err = %s", err.Error())
		}
	}

}

func ProtobufHandle(request ziface.IRequest) {
	if v, err := request.GetConnection().GetProperty("uniqueId"); err == nil {
		gateTologicMsg := logic.GateTologicMsg{
			From:      v.(int64),
			ServerId:  core.GlobalMgr.ServerId,
			Serialize: logic.GateTologicMsg_Protobuf,
			Data:      request.GetData(),
		}

		data, _ := gateTologicMsg.Marshal()

		_, _, err := global.SRPC.RPC(servivePath, serviceMethod, data, protocol.ProtoBuffer, true)
		if err != nil {
			//如果logic服务不可用，将返回err
			global.LOG.Errorf("JsonHandle RPC err = %s", err.Error())
		}
	}
}
