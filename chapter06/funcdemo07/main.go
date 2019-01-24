package main

import (
	"fmt"
	"strings"
)

// 1) 编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg),并返回一个闭包
// 2) 调用闭包，可以穿入一个文件名，如果该文件名没有指定的后缀(比如.jpg),则返回文件名.jpg,如果有指定的后缀，返回原文件名。
// 3) 要求使用闭包的方式做
// 4) strings.HasSuffix, 这个方法可以判断某个字符串是否有指定的后缀。

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 如果 有指定的后缀，返回原名字。否则添加后缀并返回。
		if strings.HasSuffix(name, suffix) == true {
			return name
		}

		return name + "." + suffix
	}
}

func main() {
	// 闭包函数 测试案例
	func1 := makeSuffix("jpg")

	fmt.Println("res1 = ", func1("helloworld"))
	fmt.Println("res2 = ", func1("今天好开心.jpg"))

	// 闭包函数的好处。
	// 如过使用传统的方法也可以是轻松实现，但是传统方法需要每次都传入后缀名，比如.jpg，
	// 而闭包 因为可以保留上一次引用的某个值，所以我们传入一次就可以反复使用。
}
