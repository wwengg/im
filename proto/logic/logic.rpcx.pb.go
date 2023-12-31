// Code generated by protoc-gen-rpcx. DO NOT EDIT.
// versions:
// - protoc-gen-rpcx v0.3.0
// - protoc          v3.20.3
// source: logic/logic.proto

package logic

import (
	context "context"
	client "github.com/smallnest/rpcx/client"
	protocol "github.com/smallnest/rpcx/protocol"
	server "github.com/smallnest/rpcx/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = context.TODO
var _ = server.NewServer
var _ = client.NewClient
var _ = protocol.NewMessage

//================== interface skeleton ===================
type LogicAble interface {
	// LogicAble can be used for interface verification.

	// GateToLogic is server rpc method as defined
	GateToLogic(ctx context.Context, args *GateTologicMsg, reply *Empty) (err error)

	// ConnectionBegin is server rpc method as defined
	ConnectionBegin(ctx context.Context, args *GateTologicMsg, reply *Empty) (err error)

	// ConnectionLost is server rpc method as defined
	ConnectionLost(ctx context.Context, args *GateTologicMsg, reply *Empty) (err error)
}

//================== server skeleton ===================
type LogicImpl struct{}

// ServeForLogic starts a server only registers one service.
// You can register more services and only start one server.
// It blocks until the application exits.
func ServeForLogic(addr string) error {
	s := server.NewServer()
	s.RegisterName("Logic", new(LogicImpl), "")
	return s.Serve("tcp", addr)
}

// GateToLogic is server rpc method as defined
func (s *LogicImpl) GateToLogic(ctx context.Context, args *GateTologicMsg, reply *Empty) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = Empty{}

	return nil
}

// ConnectionBegin is server rpc method as defined
func (s *LogicImpl) ConnectionBegin(ctx context.Context, args *GateTologicMsg, reply *Empty) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = Empty{}

	return nil
}

// ConnectionLost is server rpc method as defined
func (s *LogicImpl) ConnectionLost(ctx context.Context, args *GateTologicMsg, reply *Empty) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = Empty{}

	return nil
}

//================== client stub ===================
// Logic is a client wrapped XClient.
type LogicClient struct {
	xclient client.XClient
}

// NewLogicClient wraps a XClient as LogicClient.
// You can pass a shared XClient object created by NewXClientForLogic.
func NewLogicClient(xclient client.XClient) *LogicClient {
	return &LogicClient{xclient: xclient}
}

// NewXClientForLogic creates a XClient.
// You can configure this client with more options such as etcd registry, serialize type, select algorithm and fail mode.
func NewXClientForLogic(addr string) (client.XClient, error) {
	d, err := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	if err != nil {
		return nil, err
	}

	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer

	xclient := client.NewXClient("Logic", client.Failtry, client.RoundRobin, d, opt)

	return xclient, nil
}

// GateToLogic is client rpc method as defined
func (c *LogicClient) GateToLogic(ctx context.Context, args *GateTologicMsg) (reply *Empty, err error) {
	reply = &Empty{}
	err = c.xclient.Call(ctx, "GateToLogic", args, reply)
	return reply, err
}

// ConnectionBegin is client rpc method as defined
func (c *LogicClient) ConnectionBegin(ctx context.Context, args *GateTologicMsg) (reply *Empty, err error) {
	reply = &Empty{}
	err = c.xclient.Call(ctx, "ConnectionBegin", args, reply)
	return reply, err
}

// ConnectionLost is client rpc method as defined
func (c *LogicClient) ConnectionLost(ctx context.Context, args *GateTologicMsg) (reply *Empty, err error) {
	reply = &Empty{}
	err = c.xclient.Call(ctx, "ConnectionLost", args, reply)
	return reply, err
}

//================== oneclient stub ===================
// LogicOneClient is a client wrapped oneClient.
type LogicOneClient struct {
	serviceName string
	oneclient   *client.OneClient
}

// NewLogicOneClient wraps a OneClient as LogicOneClient.
// You can pass a shared OneClient object created by NewOneClientForLogic.
func NewLogicOneClient(oneclient *client.OneClient) *LogicOneClient {
	return &LogicOneClient{
		serviceName: "Logic",
		oneclient:   oneclient,
	}
}

// ======================================================

// GateToLogic is client rpc method as defined
func (c *LogicOneClient) GateToLogic(ctx context.Context, args *GateTologicMsg) (reply *Empty, err error) {
	reply = &Empty{}
	err = c.oneclient.Call(ctx, c.serviceName, "GateToLogic", args, reply)
	return reply, err
}

// ConnectionBegin is client rpc method as defined
func (c *LogicOneClient) ConnectionBegin(ctx context.Context, args *GateTologicMsg) (reply *Empty, err error) {
	reply = &Empty{}
	err = c.oneclient.Call(ctx, c.serviceName, "ConnectionBegin", args, reply)
	return reply, err
}

// ConnectionLost is client rpc method as defined
func (c *LogicOneClient) ConnectionLost(ctx context.Context, args *GateTologicMsg) (reply *Empty, err error) {
	reply = &Empty{}
	err = c.oneclient.Call(ctx, c.serviceName, "ConnectionLost", args, reply)
	return reply, err
}
