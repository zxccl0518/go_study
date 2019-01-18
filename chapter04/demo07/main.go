package main

import "fmt"

func main() {
	// 演示 & *

	a := 100
	fmt.Printf("a 的地址：%d", &a)

	var ptr = &a
	fmt.Printf("ptr 指针指向的值是: %v \n", *ptr)

	var n int
	var i int = 10
	var j int = 12
	// 传统的三元运算。
	// n = i < j ? i : j
	// 在go中 用if else表示
	if i > j {
		n = i
	} else {
		n = j
	}
	fmt.Println(" n = ", n)
}
