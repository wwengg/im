package middleware

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/protocol"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/httpgate/model/request"
	"github.com/wwengg/im/internal/httpgate/model/response"
	"github.com/wwengg/im/proto/httpgate"
	"github.com/wwengg/im/proto/pbcommon"
	"github.com/wwengg/simple/core/plugin"
	"go.uber.org/zap"
)

func V2Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var md map[string]string
		if md2, isExist := c.Get("JaegerMd"); isExist {
			md = md2.(map[string]string)
		}
		if span, md, err := plugin.GenerateSpanWithMap(md, "V2Handler"); err == nil {
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
		if !(isJson || isProtobuf) {
			c.AbortWithStatus(http.StatusForbidden)
			return
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
			requestJson := request.RequestJson{}
			//将前端json格式数据与LoginForm对象绑定
			err := c.BindJSON(&requestJson)
			if err != nil {
				global.LOG.Errorf("c.BindJSON:%s", err.Error())
				response.GatewayResult(pbcommon.EnumCode_Invalid, "非法参数", c)
				return
			}
			if bytes, err := json.Marshal(requestJson.Data); err == nil {
				req.Payload = bytes
			} else {
				global.LOG.Errorf("json.Marshal:%s,data:%v", err.Error(), requestJson.Data)
				response.GatewayResult(pbcommon.EnumCode_Invalid, "非法参数", c)
				return
			}
			global.LOG.Info("V2Handler Success", zap.Any("requestJson", requestJson))
		} else if isProtobuf {
			req.SetSerializeType(protocol.ProtoBuffer)
			payload, err := io.ReadAll(c.Request.Body)
			defer c.Request.Body.Close()
			if err != nil {
				global.LOG.Errorf("requestProto.Unmarshal:%s", err.Error())
				response.GatewayResult(pbcommon.EnumCode_Invalid, "非法参数", c)
				return
			}
			var requestProto httpgate.HttpRequest

			if err = requestProto.Unmarshal(payload); err != nil {
				global.LOG.Errorf("requestProto.Unmarshal:%s", err.Error())
				response.GatewayResult(pbcommon.EnumCode_Invalid, "非法参数", c)
				return
			}
			req.Payload = requestProto.Data
			global.LOG.Info("V2Handler Success", zap.Any("requestProto", requestProto))
		}
		c.Set("rpcxMessage", req)
	}
}
