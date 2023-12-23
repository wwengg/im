// @Title
// @Description
// @Author  Wangwengang  2023/12/22 08:49
// @Update  Wangwengang  2023/12/22 08:49
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wwengg/im/internal/httpgate/controller"
)

func InitSRPCRouter(Router *gin.RouterGroup) {
	Http2SRPCRouter := Router
	{
		Http2SRPCRouter.POST(":servicePath/:serviceMethod", controller.Http2RpcxPost)
	}
}
