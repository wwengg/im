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
	// objVal := reflect.ValueOf(c).Elem()
	// objType := reflect.TypeOf(*c)

	fmt.Println("===== IM Global Config =====")
	// fmt.Printf("%+v\n", *c)
	show(reflect.ValueOf(c).Elem(), reflect.TypeOf(*c))
	// for i := 0; i < objVal.NumField(); i++ {
	// 	field := objVal.Field(i)
	// 	typeField := objType.Field(i)

	// 	fmt.Printf("%s: %v\n", typeField.Name, field.Interface())
	// 	// objVal2 := field
	// 	// objType2 := field.Type()

	// 	// for a := 0; a < objVal2.NumField(); a++ {
	// 	// 	field2 := objVal2.Field(a)
	// 	// 	typeField2 := objType2.Field(a)
	// 	// 	fmt.Printf("%s: %v\n", typeField2.Name, field2.Interface())
	// 	// }
	// }
	fmt.Println("==============================")
}

func show(objVal reflect.Value, objType reflect.Type) {

	if objVal.Kind() == reflect.Slice {
		objVal.Len()
		for i := 0; i < objVal.Len(); i++ {
			b := objVal.Index(i).Kind()
			if b == reflect.Struct || b == reflect.Slice {
				show(objVal.Index(i), objVal.Index(i).Type())
			}
		}
		return
	}
	for i := 0; i < objVal.NumField(); i++ {
		field := objVal.Field(i)
		typeField := objType.Field(i)

		fmt.Printf("%s: %v\n", typeField.Name, field.Interface())
		b := field.Kind()
		if b == reflect.Struct || b == reflect.Slice {
			fmt.Printf("===== %s Desc =====\n", typeField.Name)
			show(field, field.Type())
			fmt.Println("==============================")
			fmt.Println("\n")
		}
		// if field.NumField() > 1 {
		// 	show(field, field.Type())
		// }
	}
}
