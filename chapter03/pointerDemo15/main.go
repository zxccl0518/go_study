package main

import "fmt"

func main() {
	//基本数据类型在内存布局
	var i = 10
	// i 的地址是什么呢
	fmt.Println("i 的内存地址是：", &i)

	// 下面的var ptr *int = &i
	// 1.ptr 是一个指针变量。
	// 2.ptr 的类型 *int
	// 3.ptr 本身的值&i
	var ptr *int = &i
	fmt.Printf(" ptr = %v \n", ptr)
}
