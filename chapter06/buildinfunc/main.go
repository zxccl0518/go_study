package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("num1 类型:%T, 值:%v, 地址:%v\n", num1, num1, &num1)

	// new 是一个内置函数。 内置函数还有 len() new()  make()
	num2 := new(int) // *int
	// num2 的类型%T => *int
	// num2 的值%v => 存储的内容是 地址， 这个地址指向的空间是 0
	// num2 的地址%v => 存储空间的地址
	fmt.Printf("num2 类型:%T, 值:%v, 地址:%v\n", num2, num2, &num2)
	fmt.Printf("*num2 类型:%T, 值:%v, 地址:%v\n", *num2, *num2, num2)
	// 视频138 有解释的图
	// 其实 num2 就是个int类型的指针
}
