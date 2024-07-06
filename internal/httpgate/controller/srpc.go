// @Title
// @Description
// @Author  Wangwengang  2023/12/22 08:41
// @Update  Wangwengang  2023/12/22 08:41
package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/share"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/httpgate/model/response"
	"github.com/wwengg/im/proto/pbcommon"
	"github.com/wwengg/simple/core/plugin"
	"go.uber.org/zap"
)

func Http2RpcxPost(c *gin.Context) {
	var md map[string]string
	if md2, isExist := c.Get("JaegerMd"); isExist {
		md = md2.(map[string]string)
	}
	if span, md, err := plugin.GenerateSpanWithMap(md, "V1Handler"); err == nil {
		c.Set("JaegerMd", md)
		defer span.Finish()
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, share.ReqMetaDataKey, md)
	if span, ctx2, err := plugin.GenerateSpanWithContext(ctx, fmt.Sprintf("Http2RpcxPost")); err == nil {
		ctx = ctx2
		defer span.Finish()
	}

	servicePath, isExist := c.Get("servicePath")
	if !isExist {
		global.LOG.Error("Get servicePath error")
		response.GatewayResult(pbcommon.EnumCode_Internal, "Internal Error", c)
		return
	}

	xc, err := global.SRPC.GetXClient(servicePath.(string))
	if err != nil {
		global.LOG.Errorf("GetXClient error: %v", err)
		response.GatewayResult(pbcommon.EnumCode_Internal, "Internal Error", c)
		return
	}

	req, isExist := c.Get("rpcxMessage")
	if !isExist {
		global.LOG.Error("Get rpcxClient error")
		response.GatewayResult(pbcommon.EnumCode_Internal, "Internal Error", c)
		return
	}

	respMeta, resp, err := xc.SendRaw(ctx, req.(*protocol.Message))
	global.LOG.Info("请求结束", zap.Any("respMeta", respMeta), zap.Any("resp", resp), zap.Any("err", err))
	if err != nil {
		global.LOG.Debug("code != 0", zap.String("code", err.Error()))
		if err == client.ErrXClientNoServer {
			global.LOG.Errorf("not found any server servername = %s", servicePath)
			response.GatewayResult(pbcommon.EnumCode_Internal, "err", c)
			return
		}
		if code, err := strconv.Atoi(err.Error()); err != nil {
			response.GatewayResult(pbcommon.EnumCode(code), "", c)
			return
		} else {
			global.LOG.Errorf("SRPC SendRaw error=%s", err.Error())
			response.GatewayResult(pbcommon.EnumCode_Internal, "error", c)
			return
		}
	}
	if respMeta["X-RPCX-MessageStatusType"] == "Error" {
		global.LOG.Errorf("RPCX error = %s", respMeta["X-RPCX-ErrorMessage"])
		response.GatewayResult(pbcommon.EnumCode_Internal, respMeta["X-RPCX-ErrorMessage"], c)
		return
	}

	response.SrpcResult(resp, c)
}
