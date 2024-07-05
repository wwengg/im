package service

import (
	"context"
	"encoding/json"

	"github.com/gogo/protobuf/proto"
	"github.com/smallnest/rpcx/protocol"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/logic/core"
	"github.com/wwengg/im/proto/logic"
	"go.uber.org/zap"
)

type LogicService struct {
}

// GateToLogic is server rpc method as defined
func (s *LogicService) GateToLogic(ctx context.Context, args *logic.GateTologicMsg, reply *logic.Empty) (err error) {
	global.LOG.Info("GateToLogic enter", zap.Any("args", args))
	defer global.LOG.Info("GateToLogic End", zap.Any("args", args))
	*reply = logic.Empty{}
	// 获取连接信息
	conn, err := core.GlobalMgr.GetConn(*args)
	if err != nil {
		global.LOG.Error("GlobalMgr.GetConn error", zap.Error(err))
		return nil
	}
	global.LOG.Info("GET CONN SUCCESS", zap.Any("conn", conn))
	var data logic.LogicData
	var serialize protocol.SerializeType
	switch args.Serialize {
	case logic.GateTologicMsg_None:
		return nil
	case logic.GateTologicMsg_Json:
		if err := json.Unmarshal(args.Data, &data); err != nil {
			global.LOG.Errorf("json.Unmarshal err = %s", err.Error())
			return nil
		}
		serialize = protocol.JSON
	case logic.GateTologicMsg_Protobuf:
		if err := proto.Unmarshal(args.Data, &data); err != nil {
			global.LOG.Errorf("proto.Unmarshal err = %s", err.Error())
			return nil
		}
		serialize = protocol.ProtoBuffer
	}
	if cmdService, err := GlobalCmdService.getCmdService(int32(data.Cmd)); err == nil {
		_, _, err := global.SRPC.RPC(ctx, cmdService.ServicePath, cmdService.ServiceName, data.Data, serialize, true)
		if err != nil {
			global.LOG.Error("GateToLogic RPC error", zap.Error(err))
		}
	} else {
		// TODO 将消息扔到nsq,避免消息丢失
		global.LOG.Error("GateToLogic getCmdService error", zap.Error(err))
	}
	return nil
}

// ConnectionBegin is server rpc method as defined
func (s *LogicService) ConnectionBegin(ctx context.Context, args *logic.GateTologicMsg, reply *logic.Empty) (err error) {
	global.LOG.Info("ConnectionBegin enter", zap.Any("args", args))
	defer global.LOG.Info("ConnectionBegin End", zap.Any("args", args))
	if err = core.GlobalMgr.NewConn(*args); err != nil {
		global.LOG.Error("ConnectionBegin err", zap.Error(err))
	}
	*reply = logic.Empty{}

	return nil
}

// ConnectionLost is server rpc method as defined
func (s *LogicService) ConnectionLost(ctx context.Context, args *logic.GateTologicMsg, reply *logic.Empty) (err error) {
	global.LOG.Info("ConnectionLost enter", zap.Any("args", args))
	defer global.LOG.Info("ConnectionLost End", zap.Any("args", args))
	_ = core.GlobalMgr.RemoveConn(*args)
	*reply = logic.Empty{}

	return nil
}
