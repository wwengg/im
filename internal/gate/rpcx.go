// @Title
// @Description
// @Author  Wangwengang  2023/12/22 18:53
// @Update  Wangwengang  2023/12/22 18:53
package gate

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/gate/core"
	"github.com/wwengg/im/internal/gate/service"
	"github.com/wwengg/im/model"
	"github.com/wwengg/simple/core/sconfig"
	srpc "github.com/wwengg/simple/core/srpc"
	"os"
)

var _ = srpc.TODO

func initRpcxService(serverName string, rpc sconfig.RPC, rpcService sconfig.RpcService) {
	info := model.ServerInfo{
		Name:       serverName,
		Ip:         global.CONFIG.RpcService.ServiceAddr,
		Port:       global.CONFIG.RpcService.Port,
		ServerType: model.ServerType_GateTcp,
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

	s.RegisterName(fmt.Sprintf("Gate%d", core.GlobalMgr.ServerId), new(service.Gate), "")
	s.RegisterName(fmt.Sprintf("Gate", core.GlobalMgr.ServerId), new(service.Gate), "") // 广播消息时使用
	// 执行上次程序退出未处理完的数据

	// go 协程启动rpc服务 开始接客
	go global.LOG.Error(s.Serve("tcp", fmt.Sprintf("%s:%s", global.CONFIG.RpcService.ServiceAddr, global.CONFIG.RpcService.Port)).Error())
}
