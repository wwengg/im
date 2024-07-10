package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/share"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/httpgate/model/response"
	"github.com/wwengg/im/proto/pbauth"
	"github.com/wwengg/im/proto/pbcommon"
	"github.com/wwengg/simple/core/plugin"
	"go.uber.org/zap"
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
			// 交给Casbin来拦截
			c.Next()
			return
		}
		appId, isExist := c.Get("appId")
		if !isExist {
			response.GatewayResult(pbcommon.EnumCode_ParamError, "ParamError", c)
			c.Abort()
			return
		}

		ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, md)
		args := pbauth.ParseTokenArgs{Token: token.(string)}
		payload, _ := args.Marshal()
		_, resp, _ := global.SRPC.RPCProtobuf(ctx, fmt.Sprintf("Auth%d", appId.(int64)), "ParseToken", payload)
		var reply pbauth.ParseTokenReply
		_ = reply.Unmarshal(resp)
		if reply.Code == pbcommon.EnumCode_Success {
			c.Set("appId", reply.AppId)
			c.Set("userId", reply.Id)
			c.Set("roleId", reply.RoleId)
			if reply.NewToken != "" {
				c.Header("new-token", reply.NewToken)
			}
			c.Next()
		}
		global.LOG.Info("JwtHandler:", zap.Any("reply", reply))
	}
}
