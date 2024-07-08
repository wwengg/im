package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wwengg/im/internal/httpgate/model/response"
	"github.com/wwengg/im/proto/pbcommon"
	"github.com/wwengg/simple/core/plugin"
)

// c.Get("JaegerMd") && c.Get("tokenData") && c.Get("appId")
func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var md map[string]string
		if md2, isExist := c.Get("JaegerMd"); isExist {
			md = md2.(map[string]string)
		}
		if span, md3, err := plugin.GenerateSpanWithMap(md, "JwtHandler"); err == nil {
			md = md3
			c.Set("JaegerMd", md3)
			defer span.Finish()
		}

		token, isExist := c.Get("tokenData")
		if !isExist {
			response.GatewayResult(pbcommon.EnumCode_Internal, "InternalError", c)
			return
		}
		appId, isExist := c.Get("appId")
		if !isExist {
			response.GatewayResult(pbcommon.EnumCode_Internal, "InternalError", c)
			return
		}

	}
}
