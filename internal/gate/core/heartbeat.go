// @Title
// @Description
// @Author  Wangwengang  2024/2/6 11:02
// @Update  Wangwengang  2024/2/6 11:02
package core

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// User-defined heartbeat message processing method
// 用户自定义的心跳检测消息处理方法
func MyHeartBeatMsg(conn ziface.IConnection) []byte {
	return []byte("heartbeat, I am server, I am alive")
}

// User-defined handling method for remote connection not alive.
// 用户自定义的远程连接不存活时的处理方法
func MyOnRemoteNotAlive(conn ziface.IConnection) {
	fmt.Println("myOnRemoteNotAlive is Called, connID=", conn.GetConnID(), "remoteAddr = ", conn.RemoteAddr())
	//关闭链接
	conn.Stop()
}

// User-defined method for handling heartbeat messages (用户自定义的心跳检测消息处理方法)
type MyHeartBeatRouter struct {
	znet.BaseRouter
}

func (r *MyHeartBeatRouter) Handle(request ziface.IRequest) {
	fmt.Println("in MyHeartBeatRouter Handle, recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}
