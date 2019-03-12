package exercise

import (
	"reflect"
	"testing"
)

type user struct {
	Userid string
	Name   string
}

func TestReflectStruct(t *testing.T) {
	// 使用反射 操作任何结构体的案例。
	var (
		model *user
		sv    reflect.Value
	)

	model = &user{}

	sv = reflect.ValueOf(model)
	sv = sv.Elem()

	sv.FieldByName("Userid").SetString("反射演示 操作结构体， 设置 Userid 字段的内容。")
	sv.FieldByName("Name").SetString("反射演示 操作结构体， 设置 Name 字段的内容。")

	t.Log("model = ", model)

}
