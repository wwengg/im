package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wwengg/im/global"
	"github.com/wwengg/simple/core/plugin"
)

func RPCXHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		servicePath := c.Param("servicePath")

		if servicePath == "" {
			global.LOG.Error("empty servicepath")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		serviceMethod := c.Param("serviceMethod")
		if serviceMethod == "" {
			global.LOG.Error("empty servicemethod")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// 首字母转大写
		servicePath = strings.ToUpper(servicePath[:1]) + servicePath[1:]
		serviceMethod = strings.ToUpper(serviceMethod[:1]) + serviceMethod[1:]

		var md map[string]string
		if md2, isExist := c.Get("JaegerMd"); isExist {
			md = md2.(map[string]string)
		}
		if span, md, err := plugin.GenerateSpanWithMap(md, fmt.Sprintf("RPCXHandler:%s.%s", servicePath, serviceMethod)); err == nil {
			c.Set("JaegerMd", md)
			defer span.Finish()
		}

		req := global.SRPC.GetReq(servicePath, serviceMethod)
		c.Set("rpcxMessage", req)

		c.Set("servicePath", servicePath)
	}
}
