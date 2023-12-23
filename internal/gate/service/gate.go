// @Title
// @Description
// @Author  Wangwengang  2023/12/22 18:53
// @Update  Wangwengang  2023/12/22 18:53
package service

import (
	"context"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/gate/service/impl"
	"github.com/wwengg/im/proto/pbgate"
	"go.uber.org/zap"
)

type Gate struct {
}

func (s *Gate) SendMsg(ctx context.Context, args *pbgate.SendMsgData, reply *pbgate.SendMsgDataReply) error {
	global.LOG.Info("SendMsg enter", zap.Any("SendMsgData", args))
	impl.SendMsg(args)
	reply.Code = pbgate.SendMsgDataReply_Success
	return nil
}
