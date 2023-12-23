// @Title
// @Description
// @Author  Wangwengang  2023/12/22 08:43
// @Update  Wangwengang  2023/12/22 08:43
package initialize

import (
	"github.com/wwengg/im/global"
	"github.com/wwengg/simple/core/slog"
)

func InitSlog() {
	// 初始化日志
	global.LOG = slog.NewZapLog(&global.CONFIG.Slog)
}
