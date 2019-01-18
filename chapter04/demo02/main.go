package main

import "fmt"

func main() {
	// 在golang中， ++ -- 只能独立使用。
	var i int = 9
	var a int
	// a = i++ // 错误，i++只能独立使用。 // syntax error 语法错误，只允许独立使用。
	// a = i-- // 错误，i--只能独立使用。 // syntax error 语法错误，只允许独立使用。

	// 正确使用方法。
	i++
	a = i
	fmt.Println("i ++ :", a)
	i--
	a = i
	fmt.Println("i -- :", a)

	// 错误的使用方法。
	// if i++ > 0 {
	// 	fmt.Println("ok")
	// }

	// golang 的 ++ 、-- 智能卸载变量后面，不能写在变量的前面，
	var n int = 10
	n++
	// ++n  ,报错，语法错误。 错误的写法。
	fmt.Println("n = ", n)
}
