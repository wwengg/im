// @Title
// @Description
// @Author  Wangwengang  2023/12/22 11:09
// @Update  Wangwengang  2023/12/22 11:09
package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/wwengg/im/internal/gate/controller"
	"github.com/wwengg/im/proto/pbgate"
)

func InitGate2LogicRouter(group ziface.IGroupRouterSlices) {
	group.AddHandler(uint32(pbgate.Cmd_JSON), controller.JsonHandle)
	group.AddHandler(uint32(pbgate.Cmd_ProtoBuffer), controller.ProtobufHandle)
}
