package main

import "fmt"

// 演示golang中 + 的使用方式。
func main() {
	var i = 1
	var j = 2
	var r = i + j
	fmt.Println("r = ", r)

	var str1 = "hello"
	var str2 = "world"
	var res = str1 + str2
	fmt.Println("res = ", res)

	// 经过验证，一下2中写法都是可以的，自动判断变量类型
	stringValue_1 := "想知道这种写法是不是自动类型解析。 := 1"
	fmt.Println("value = ", stringValue_1)

	// 结论是，这种写法的方式也是可行的，但是不推荐。
	// var stringValue_2 = "看看这个 value2的 写法方式 是不是可行。"
	var stringValue_2 string = "看看这个 value2的 写法方式 是不是可行。"
	fmt.Println("value = ", stringValue_2)
}
