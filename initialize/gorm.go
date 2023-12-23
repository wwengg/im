// @Title
// @Description
// @Author  Wangwengang  2023/12/21 23:58
// @Update  Wangwengang  2023/12/21 23:58
package initialize

import (
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/model"
	"github.com/wwengg/simple/core/store"
)

func InitDB() {
	// 初始化DBList
	global.DBList = store.DBList(&global.CONFIG.DBList)

	global.DBUpms = global.DBList["upms"]

	// 创建初始化upms的数据库表
	global.DBUpms.AutoMigrate(
		model.ServerInfo{},
	)
}
