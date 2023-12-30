// @Title
// @Description
// @Author  Wangwengang  2023/12/21 22:35
// @Update  Wangwengang  2023/12/21 22:35
package global

import (
	"fmt"
	"github.com/wwengg/simple/core/sconfig"
	"reflect"
)

type Config struct {
	sconfig.Config `yaml:",inline" mapstructure:",squash"`

	Zinx Zinx `json:"zinx" yaml:"zinx" mapstructure:"zinx"`
}

func (c *Config) Show() {
	objVal := reflect.ValueOf(c).Elem()
	objType := reflect.TypeOf(*c)

	fmt.Println("===== IM Global Config =====")
	for i := 0; i < objVal.NumField(); i++ {
		field := objVal.Field(i)
		typeField := objType.Field(i)

		fmt.Printf("%s: %v\n", typeField.Name, field.Interface())
	}
	fmt.Println("==============================")
}
