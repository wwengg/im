// @Title
// @Description
// @Author  Wangwengang  2023/12/22 08:41
// @Update  Wangwengang  2023/12/22 08:41
package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/httpgate/model/request"
	"github.com/wwengg/im/internal/httpgate/model/response"
	"github.com/wwengg/im/proto/httpgate"
	"github.com/wwengg/im/proto/pbcommon"
	"github.com/wwengg/simple/core/plugin"
	"go.uber.org/zap"
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
		response.GatewayResult(pbcommon.EnumCode_Invalid, "非法参数", c)
		return
	}

	serviceMethod := c.Param("serviceMethod")
	if serviceMethod == "" {
		global.LOG.Error("empty servicemethod")
		response.GatewayResult(pbcommon.EnumCode_Invalid, "非法参数", c)
		return
	}

	// 首字母小写转大写
	servicePath = strings.ToUpper(servicePath[:1]) + servicePath[1:]
	serviceMethod = strings.ToUpper(serviceMethod[:1]) + serviceMethod[1:]
	global.LOG.Info("请求开始", zap.Any("servicePath", servicePath), zap.Any("serviceMethod", serviceMethod))
	if span, _, err := plugin.GenerateSpanWithContext(c, fmt.Sprintf("Http2RpcxPost:%s.%s", servicePath, serviceMethod)); err == nil {
		defer span.Finish()
	}
	var err error
	meta := make(map[string]string, 0)
	var resp []byte
	if isJson {
		if servicePath == "DeviceReport" {
			payload, err := io.ReadAll(c.Request.Body)
			defer c.Request.Body.Close()
			if err != nil {
				global.LOG.Error(err.Error())
			}
			global.LOG.Infof("DeviceReport payload: %s", string(payload))
			meta, resp, err = global.SRPC.RPCJson(c, servicePath, serviceMethod, payload)
		} else {
			requestJson := request.RequestJson{}
			//将前端json格式数据与LoginForm对象绑定
			err := c.BindJSON(&requestJson)
			if err != nil {
				response.GatewayResult(pbcommon.EnumCode_Invalid, "非法参数", c)
				return
			}
			if bytes, err := json.Marshal(requestJson.Data); err == nil {
				meta, resp, err = global.SRPC.RPCJson(c, servicePath, serviceMethod, bytes)
			} else {
				response.GatewayResult(pbcommon.EnumCode_Invalid, "非法参数", c)
				return
			}
		}
	} else if isProtobuf {
		payload, err := io.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()
		if err != nil {
			global.LOG.Error(err.Error())
		}
		var requestProto httpgate.HttpRequest

		if err = requestProto.Unmarshal(payload); err != nil {
			global.LOG.Error(err.Error())
			response.GatewayResult(pbcommon.EnumCode_Invalid, "非法参数", c)
			return
		}
		global.LOG.Info("SendRaw", zap.String("servicePath", servicePath), zap.String("serviceMethod", serviceMethod), zap.Any("requestProto", requestProto))
		meta, resp, err = global.SRPC.RPCProtobuf(c, servicePath, serviceMethod, requestProto.Data)
	}
	global.LOG.Info("请求结束", zap.Any("meta", meta), zap.Any("resp", resp), zap.Any("err", err))
	if err != nil {
		global.LOG.Debug("code != 0", zap.String("code", err.Error()))
		if err == client.ErrXClientNoServer {
			global.LOG.Errorf("not found any server servername = %s,method = %s", servicePath, serviceMethod)
			response.GatewayResult(pbcommon.EnumCode_Internal, "err", c)
			return
		}
		if code, err := strconv.Atoi(err.Error()); err != nil {
			response.GatewayResult(pbcommon.EnumCode(code), "error", c)
			return
		} else {
			global.LOG.Errorf("SRPC SendRaw error=%s", err.Error())
			response.GatewayResult(pbcommon.EnumCode_Internal, "error", c)
			return
		}
	}
	if meta["X-RPCX-MessageStatusType"] == "Error" {
		global.LOG.Errorf("RPCX error = %s", meta["X-RPCX-ErrorMessage"])
		response.GatewayResult(pbcommon.EnumCode_Internal, meta["X-RPCX-ErrorMessage"], c)
		return
	}

	response.SrpcResult(resp, c)
}
