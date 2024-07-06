package middleware

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/protocol"
	"github.com/wwengg/im/global"
	"github.com/wwengg/simple/core/plugin"
)

func V1Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var md map[string]string
		if md2, isExist := c.Get("JaegerMd"); isExist {
			md = md2.(map[string]string)
		}
		if span, md, err := plugin.GenerateSpanWithMap(md, "V1Handler"); err == nil {
			c.Set("JaegerMd", md)
			defer span.Finish()
		}

		var isJson, isProtobuf bool
		if a, exist := c.Get("isJson"); exist {
			isJson = a.(bool)
		}
		if a, exist := c.Get("isProtobuf"); exist {
			isProtobuf = a.(bool)
		}

		var req *protocol.Message
		if req2, isExist := c.Get("rpcxMessage"); !isExist {
			c.AbortWithStatus(http.StatusForbidden)
			return
		} else {
			req = req2.(*protocol.Message)
		}

		if isJson {
			req.SetSerializeType(protocol.JSON)
		} else if isProtobuf {
			req.SetSerializeType(protocol.ProtoBuffer)
		} else {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		payload, err := io.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()
		if err != nil {
			global.LOG.Error(err.Error())
		}
		if isJson {
			global.LOG.Infof("json payload: %s", string(payload))
		}
		c.Set("rpcxMessage", req)
	}
}
