package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/share"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/httpgate/model/response"
	"github.com/wwengg/im/proto/pbcasbinRule"
	"github.com/wwengg/im/proto/pbcommon"
	"github.com/wwengg/simple/core/plugin"
)

// c.Get("JaegerMd") && c.Get("roleId")
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var md map[string]string
		if md2, isExist := c.Get("JaegerMd"); isExist {
			md = md2.(map[string]string)
		}
		if span, md3, err := plugin.GenerateSpanWithMap(md, "CasbinHandler"); err == nil {
			md = md3
			c.Set("JaegerMd", md3)
			defer span.Finish()
		}

		// 获取请求的URI
		path := c.Request.URL.RequestURI()
		// 获取请求方法
		method := c.Request.Method

		a, exist := c.Get("roleId")
		if exist {
			a = a.(int64)
		} else {
			a = int64(0) // 游客
		}
		args := pbcasbinRule.CheckCasbinRuleArgs{
			Role:   a.(int64),
			Path:   path,
			Method: method,
		}

		ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, md)
		payload, _ := args.Marshal()
		_, resp, _ := global.SRPC.RPCProtobuf(ctx, "CasbinRule", "CheckCasbinRule", payload)
		var reply pbcasbinRule.CheckCasbinRuleReply
		_ = reply.Unmarshal(resp)
		if reply.Ok {
			return
		} else {
			response.GatewayResult(pbcommon.EnumCode_Forbidden, "权限不足", c)
			c.Abort()
		}
		global.LOG.Infof("CasbinHandler path: %s, method: %s", path, method)
	}
}
