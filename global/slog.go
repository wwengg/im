// @Title
// @Description
// @Author  Wangwengang  2023/12/22 08:42
// @Update  Wangwengang  2023/12/22 08:42
package global

import "github.com/wwengg/simple/core/slog"

var LOG slog.Slog

func InitSlog() {
	// 初始化日志
	LOG = slog.NewZapLog(&CONFIG.Slog)
}
