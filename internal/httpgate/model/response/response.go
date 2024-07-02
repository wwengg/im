// @Title
// @Description
// @Author  Wangwengang  2023/12/22 09:37
// @Update  Wangwengang  2023/12/22 09:37
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/proto/httpgate"
	"github.com/wwengg/im/proto/pbcommon"
	"go.uber.org/zap"
)

func GatewayResult(code pbcommon.EnumCode, msg string, c *gin.Context) {
	var isJson, isProtobuf bool
	if a, exist := c.Get("isJson"); exist {
		isJson = a.(bool)
	}
	if a, exist := c.Get("isProtobuf"); exist {
		isProtobuf = a.(bool)
	}
	if isJson {
		c.JSON(http.StatusOK, httpgate.HttpResponse{
			Code: code,
			Msg:  msg,
		})
	} else if isProtobuf {
		reply := &httpgate.HttpResponse{
			Code: code,
			Msg:  msg,
		}
		data, _ := reply.Marshal()
		c.Data(http.StatusOK, "application/x-protobuf", data)
	} else {

	}
}

func Ok(c *gin.Context) {
	GatewayResult(200, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	GatewayResult(200, message, c)
}

func Fail(code pbcommon.EnumCode, c *gin.Context) {
	GatewayResult(code, "error", c)
}

func SrpcResult(data []byte, c *gin.Context) {
	var isJson, isProtobuf bool
	if a, exist := c.Get("isJson"); exist {
		isJson = a.(bool)
	}
	if a, exist := c.Get("isProtobuf"); exist {
		isProtobuf = a.(bool)
	}
	if isJson {
		global.LOG.Info("SrpcResult", zap.Any("responseJson", data))
		c.Data(http.StatusOK, "application/json; charset=utf-8", data)
	} else if isProtobuf {
		var responseProto httpgate.HttpResponse
		responseProto.Unmarshal(data)
		global.LOG.Info("Return", zap.Any("responseProto", responseProto))
		c.Data(http.StatusOK, "application/x-protobuf", data)
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}
