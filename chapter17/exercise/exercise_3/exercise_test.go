package exercise

import (
	"reflect"
	"testing"
)

type User struct {
	Name   string
	UserId string
}

func TestReflectStruct(t *testing.T) {
	// 反射 演示创建结构体。
	var (
		st    reflect.Type
		sv    reflect.Value
		model *User
	)

	st = reflect.TypeOf(model)
	t.Log("类型 = ", st.Kind().String())
	st = st.Elem()
	t.Log("leixin : ", st.Kind().String())

	sv = reflect.New(st) // new 返回一个value类性质，该值持有一个指向类型为type的心申请的零值的指针。
	t.Log("类型 = ", sv.Kind().String())
	t.Log("类型 = ", sv.Elem().Kind().String())
	model = sv.Interface().(*User)

	sv = sv.Elem()
	sv.FieldByName("Name").SetString("哈哈哈哈")
	sv.FieldByName("UserId").SetString("123123123")

	t.Log("model = ", model)
}
