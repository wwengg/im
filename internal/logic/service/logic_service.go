package service

import (
	"context"
	"github.com/wwengg/im/proto/logic"
)

type LogicService struct {
}

// GateToLogic is server rpc method as defined
func (s *LogicService) GateToLogic(ctx context.Context, args *logic.GateTologicMsg, reply *logic.Empty) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = logic.Empty{}

	return nil
}

// ConnectionBegin is server rpc method as defined
func (s *LogicService) ConnectionBegin(ctx context.Context, args *logic.GateTologicMsg, reply *logic.Empty) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = logic.Empty{}

	return nil
}

// ConnectionLost is server rpc method as defined
func (s *LogicService) ConnectionLost(ctx context.Context, args *logic.GateTologicMsg, reply *logic.Empty) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = logic.Empty{}

	return nil
}
