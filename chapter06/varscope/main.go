package main

import (
	"fmt"

	gutils "github.com/zxccl0518/go_study/chapter06/globalutils"
)

// 函数外部声明/定义的变量叫做全局变量。
// 作用域在整个包都有效，如果其首字母为大写，则作用域为整个程序。
var age = 50
var Name = "jack~"

// 函数
func test() {
	// age 和 Name的作用于就只在test函数内部。
	age := 10
	Name := "tom~"
	fmt.Println("test age = ", age)
	fmt.Println("test Name = ", Name)
}

func main() {

	test()

	// 由于作用域不一样，及时同一个变量名字 也是对应的不同的值。
	fmt.Println("main age = ", age)
	fmt.Println("main Name = ", Name)

	fmt.Println("varscope包下的 变量Name是全程序都可以使用。 Name = ", gutils.GlobalValue)

	gutils.Test()
}
