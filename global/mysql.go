// @Title
// @Description
// @Author  Wangwengang  2023/12/21 23:59
// @Update  Wangwengang  2023/12/21 23:59
package global

import (
	"github.com/wwengg/im/model"
	"github.com/wwengg/simple/core/store"
	"gorm.io/gorm"
)

var (
	DBList map[string]*gorm.DB
	DBUpms *gorm.DB
)

func InitDB() {
	// 初始化DBList
	DBList = store.DBList(&CONFIG.DBList)

	DBUpms = DBList["upms"]

	// 创建初始化upms的数据库表
	DBUpms.AutoMigrate(
		model.ServerInfo{},
	)
}
