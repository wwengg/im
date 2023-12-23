// @Title
// @Description
// @Author  Wangwengang  2023/12/23 12:00
// @Update  Wangwengang  2023/12/23 12:00
package impl

import (
	"github.com/wwengg/im/internal/gate/core"
	"github.com/wwengg/im/proto/pbgate"
)

func SendMsg(args *pbgate.SendMsgData) {
	if args.Type == pbgate.SendMsgData_To_All {
		core.GlobalMgr.SendMsgToAll(args.DataJson, args.DataProto)
	} else if args.Type == pbgate.SendMsgData_To_Conn {
		core.GlobalMgr.SendMsgToConn(uint64(args.To), args.DataJson, args.DataProto)
	}
}
