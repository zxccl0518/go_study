package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// 通过反射获取的传入的变量 的type，kind，值
	//1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	rValue := reflect.ValueOf(b)
	fmt.Println("rValue = ", rValue)
	// n1 := 10
	// n2 := 2 + n1
	// n2 := 2 + rValue //报错:invalid operation: 2 + rValue (mismatched types int and reflect.Value)
	n2 := 2 + rValue.Int() // 这才是真正拿到 rValue 反射之前的变量的值
	fmt.Println("n2 = ", n2)
	fmt.Printf("rValue = %v, type = %T\n", rValue, rValue)

	// 下面 我们将rValue 转成interface{}
	iV := rValue.Interface()
	// 将 interface{} 重新断言 转成所需要的类型。
	num2 := iV.(int)
	fmt.Println(" num2 = ", num2)
}

type Student struct {
	Name string
	Age  int
}
type Monster struct {
}

// 专门演示反射[对结构体的反射。]
func reflectTest02(b interface{}) {
	// 通过反射获取的传入的变量 的type，kind，值
	//1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType, " rType = ", rType.Kind())

	// 2.获取到reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Println("struct rValue = ", rValue)
	fmt.Printf("struct rValue = %v, rValue type = %T \n", rValue, rValue)

	// 3.获取 变量对应的kind
	// 1) rValue.Kind() == >
	// 2) rType.Kind() ==>
	fmt.Printf("kind = %v, kind = %v \n", rValue.Kind(), rType.Kind())

	// 下面 我们将rValue 转成interface{}
	iV := rValue.Interface()
	fmt.Printf("iv = %v, iv type = %T\n", iV, iV)
	// 将 interface{} 重新断言 转成所需要的类型。
	// value, ok := iV.(Student)
	switch iV.(type) {
	case Student:
		fmt.Println("是 student 结构体类型。")
		fmt.Printf(" Name = %v, Age = %v\n", iV.(Student).Name, iV.(Student).Age)

	case Monster:
		fmt.Println("是 Monster 结构体类型。")
	}
}

func main() {
	// 请编写一个案例。
	// 演示对(基本数据类型，interface{}，reflect.Value进行反射的基本操作)

	// 1.先定义一个int
	// var num int = 100
	// reflectTest01(num)

	// 2. 定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}

	reflectTest02(stu)
}
