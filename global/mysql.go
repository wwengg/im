// @Title
// @Description
// @Author  Wangwengang  2023/12/21 23:59
// @Update  Wangwengang  2023/12/21 23:59
package global

import "gorm.io/gorm"

var (
	DBList map[string]*gorm.DB
	DBUpms *gorm.DB
)
