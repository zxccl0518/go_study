package main

import (
	"fmt"
	"reflect"
)

func main() {
	// str := "tom"

	// fs := reflect.ValueOf(str)
	// fs.SetString("jack")

	// fmt.Println("str = ", str)

	num := 123
	rVal := reflect.ValueOf(&num)
	fmt.Printf("rVal type = %T ,rVal.kind = %v\n", rVal, rVal.Kind())
	rVal.Elem().SetInt(999)
	fmt.Println(" num = ", num)

	// 通过反射的机制，要是想改变原来的值 就必须传入变量的地址。 传入地址就需要Elem() 方法拿到  v持有的指针指向的值的Value封装
}
