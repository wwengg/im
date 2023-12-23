// @Title
// @Description
// @Author  Wangwengang  2023/12/22 09:37
// @Update  Wangwengang  2023/12/22 09:37
package response

import (
	"github.com/gin-gonic/gin"
	"github.com/wwengg/im/proto/httpgate"
	"net/http"
)

func GatewayResult(code int32, msg string, c *gin.Context) {
	var isJson, isProtobuf bool
	if a, exist := c.Get("isJson"); exist {
		isJson = a.(bool)
	}
	if a, exist := c.Get("isProtobuf"); exist {
		isProtobuf = a.(bool)
	}
	if isJson {
		c.JSON(http.StatusOK, httpgate.Response{
			Code: code,
			Msg:  msg,
		})
	} else if isProtobuf {
		reply := &httpgate.Response{
			Code: code,
			Msg:  msg,
		}
		c.ProtoBuf(http.StatusOK, reply)

	} else {

	}
}

func Ok(c *gin.Context) {
	GatewayResult(200, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	GatewayResult(200, message, c)
}

func Fail(code int32, c *gin.Context) {
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
		c.Data(http.StatusOK, "application/json; charset=utf-8", data)
	} else if isProtobuf {
		c.Data(http.StatusOK, "application/x-protobuf", data)
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}
