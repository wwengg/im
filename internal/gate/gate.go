// @Title
// @Description
// @Author  Wangwengang  2023/12/22 10:42
// @Update  Wangwengang  2023/12/22 10:42
package gate

import (
	"fmt"
	"github.com/aceld/zinx/zdecoder"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/gate/controller"
	"github.com/wwengg/im/internal/gate/core"
	"github.com/wwengg/im/internal/gate/router"
	"github.com/wwengg/im/model"
	"os"
	"runtime"
	"strconv"
)

var GateServer ziface.IServer

func InitGateTcp() {
	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量

	info := model.ServerInfo{
		Name:       global.CONFIG.Zinx.Name,
		Ip:         global.CONFIG.Zinx.Host,
		Port:       strconv.Itoa(global.CONFIG.Zinx.TCPPort),
		ServerType: model.ServerType_GateTcp,
	}
	if err := model.ServerInit(&info); err != nil {
		global.LOG.Errorf("model.ServerInit err = %s", err.Error())
		os.Exit(1)
	}

	core.NewGlobalMgr(info.ID)

	GateServer = znet.NewUserConfDefaultRouterSlicesServer(global.CONFIG.Zinx.ToZinxConfig())

	GateServer.SetOnConnStart(controller.DoConnectionBegin)
	GateServer.SetOnConnStop(controller.DoConnectionLost)
	GateServer.AddInterceptor(zdecoder.NewHTLVCRCDecoder())

	group1 := GateServer.Group(1, 2, controller.Base)
	{
		router.InitGate2LogicRouter(group1)
	}

	GateServer.Start()
}
