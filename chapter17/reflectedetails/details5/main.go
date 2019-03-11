package main

import (
	"fmt"
	"reflect"
)

// 通过反射，修改
// num int 的值
// 修改student的值

func reflect01(b interface{}) {
	// 2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal 的类型是 %v \n", rVal.Kind()) //rVal 的类型是 ptr

	// rVal.SetInt(20) //reflect.Value.SetInt using unaddressable value

	// 因为传入的是一个指针
	rVal.Elem().SetInt(20)
	// func (v Value)SetInt(x uint64){}
	// 这个方法 操作对象是 v Value， 方法的解释是：设置v的持有值，如果v的kind 不是Int、Int16，int32，int64之一或者是v.Canset()返回假，会panic
	// 这里 传进来的是int的指针 ptr类型，而不是int类型。 所以要用这个方法的话一定要找到对应类型的实例。
	// rVal.Elem()这个方法恰恰是 拿到地址类型 对应的值
}

func main() {
	var num = 10
	reflect01(&num)
	fmt.Println("num = ", num)
}
