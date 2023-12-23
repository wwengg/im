// @Title
// @Description
// @Author  Wangwengang  2023/12/22 08:44
// @Update  Wangwengang  2023/12/22 08:44
package initialize

import (
	"github.com/wwengg/im/global"
	"github.com/wwengg/simple/core/srpc"
)

func InitSRPC() {
	// 初始化SRPC
	global.SRPC = srpc.NewSRPCClients(&global.CONFIG.RPC)
}
