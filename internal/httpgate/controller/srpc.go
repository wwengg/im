// @Title
// @Description
// @Author  Wangwengang  2023/12/22 08:41
// @Update  Wangwengang  2023/12/22 08:41
package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/httpgate/model/request"
	"github.com/wwengg/im/internal/httpgate/model/response"
	"go.uber.org/zap"
	"io"
	"strconv"
	"strings"
)

func Http2RpcxPost(c *gin.Context) {
	var isJson, isProtobuf bool
	if a, exist := c.Get("isJson"); exist {
		isJson = a.(bool)
	}
	if a, exist := c.Get("isProtobuf"); exist {
		isProtobuf = a.(bool)
	}
	servicePath := c.Param("servicePath")

	if servicePath == "" {
		global.LOG.Error("empty servicepath")
	}

	serviceMethod := c.Param("serviceMethod")
	if serviceMethod == "" {
		global.LOG.Error("empty servicemethod")
	}

	// 首字母小写转大写
	servicePath = strings.ToUpper(servicePath[:1]) + servicePath[1:]
	serviceMethod = strings.ToUpper(serviceMethod[:1]) + serviceMethod[1:]

	var err error
	meta := make(map[string]string, 0)
	var resp []byte
	if isJson {
		requestJson := request.RequestJson{}
		//将前端json格式数据与LoginForm对象绑定
		err := c.BindJSON(&requestJson)
		if err != nil {
			response.GatewayResult(403, "非法参数", c)
			return
		}
		if bytes, err := json.Marshal(requestJson.Data); err == nil {
			meta, resp, err = global.SRPC.RPCJson(servicePath, serviceMethod, bytes)
		} else {
			response.GatewayResult(403, "非法参数", c)
			return
		}
	} else if isProtobuf {
		payload, err := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		if err != nil {
			global.LOG.Error(err.Error())
		}
		meta, resp, err = global.SRPC.RPCProtobuf(servicePath, serviceMethod, payload)
	}
	global.LOG.Info("请求结束", zap.Any("meta", meta), zap.Any("resp", resp), zap.Error(err))
	if err != nil {
		global.LOG.Debug("code != 0", zap.String("code", err.Error()))
		if err == client.ErrXClientNoServer {
			global.LOG.Errorf("not found any server servername = %s,method = %s", servicePath, serviceMethod)
			response.GatewayResult(500, "err", c)
			return
		}
		if code, err := strconv.Atoi(err.Error()); err != nil {
			response.GatewayResult(int32(code), "error", c)
			return
		} else {
			global.LOG.Errorf("SRPC SendRaw error=%s", err.Error())
			response.GatewayResult(500, "error", c)
			return
		}
	}
	if meta["X-RPCX-MessageStatusType"] == "Error" {
		global.LOG.Errorf("RPCX error = %s", meta["X-RPCX-ErrorMessage"])
		response.GatewayResult(500, meta["X-RPCX-ErrorMessage"], c)
		return
	}

	response.SrpcResult(resp, c)
}
