// @Title
// @Description
// @Author  Wangwengang  2023/12/22 12:52
// @Update  Wangwengang  2023/12/22 12:52
package controller

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/wwengg/im/internal/gate/core"
	"github.com/wwengg/im/proto/pbgate"
)

func Base(request ziface.IRequest) {
	fmt.Println("Base")
	var serializeType pbgate.Cmd
	serializeType = pbgate.Cmd(request.GetMsgID())

	request.GetConnection().SetProperty("serializeType", serializeType)

}

func DoConnectionBegin(conn ziface.IConnection) {
	core.GlobalMgr.ConnectionEnterFirst(conn)
}

func DoConnectionLost(conn ziface.IConnection) {
	core.GlobalMgr.ConnectionLost(conn)
}
