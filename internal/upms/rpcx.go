package upms

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/model"
	"github.com/wwengg/simple/core/sconfig"
	"github.com/wwengg/simple/core/srpc"
	"go.uber.org/zap"
	"os"
)

var _ = srpc.TODO

func InitRpcxService(serverName string, rpc sconfig.RPC, rpcService sconfig.RpcService) {
	info := model.ServerInfo{
		Name:       serverName,
		Ip:         global.CONFIG.RpcService.ServiceAddr,
		Port:       global.CONFIG.RpcService.Port,
		ServerType: model.ServerType_RPCX_service,
	}
	if err := model.ServerInit(&info); err != nil {
		global.LOG.Errorf("model.ServerInit err = %s", err.Error())
		os.Exit(1)
	}
	// 初始化rpcx服务
	s := server.NewServer()
	// 开启rpcx监控
	s.EnableProfile = true
	// 服务注册中心
	srpc.AddRegistryPlugin(s, rpc, rpcService)

	//s.RegisterName("Logic", new(service.LogicService), "")
	// 执行上次程序退出未处理完的数据

	// go 协程启动rpc服务 开始接客
	addr := fmt.Sprintf("%s:%s", global.CONFIG.RpcService.ServiceAddr, global.CONFIG.RpcService.Port)
	global.LOG.Info("rpcx", zap.Any("rpcxservice", global.CONFIG.RpcService), zap.String("addr", addr))
	global.LOG.Error(s.Serve("tcp", addr).Error())
}
