// @Title
// @Description
// @Author  Wangwengang  2023/12/22 08:45
// @Update  Wangwengang  2023/12/22 08:45
package httpgate

import (
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/httpgate/middleware"
	"github.com/wwengg/im/internal/httpgate/router"
	"github.com/wwengg/simple/core/snet/http"
)

func InitGinHttpGate() *http.GinEngine {
	// 初始化gin
	GinEngine := http.NewGinEngine(&global.CONFIG.Gateway)

	// 配置路由，中间件
	publicGroup := GinEngine.GetPublicRouterGroup()
	privateGroup := GinEngine.GetPrivateRouterGroup()
	publicGroup.Use(middleware.BaseHandler())
	{
		router.InitSRPCRouter(publicGroup)
	}
	privateGroup.Use(middleware.BaseHandler())
	{
		router.InitSRPCRouter(privateGroup)
	}
	return GinEngine
}
