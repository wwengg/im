// @Title
// @Description
// @Author  Wangwengang  2023/12/21 22:35
// @Update  Wangwengang  2023/12/21 22:35
package global

import (
	"fmt"
	"reflect"

	"github.com/wwengg/simple/core/sconfig"
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
		objVal2 := reflect.ValueOf(objVal.Field(i)).Elem()
		objType2 := reflect.TypeOf(objVal.Field(i))

		for a := 0; a < objVal2.NumField(); a++ {
			field2 := objVal2.Field(a)
			typeField2 := objType2.Field(a)
			fmt.Printf("%s: %v\n", typeField2.Name, field2.Interface())
		}
		field := objVal.Field(i)
		typeField := objType.Field(i)

		fmt.Printf("%s: %v\n", typeField.Name, field.Interface())
	}
	fmt.Println("==============================")
}
