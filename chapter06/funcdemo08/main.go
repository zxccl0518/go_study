package main

import (
	"fmt"
)

// defer 演示函数
func sum(n1 int, n2 int) int {
	// 当执行到defer时，暂时不执行，会将deger后面的语句压入到独立的栈(defer栈)
	// 当函数执行完毕，在熊defer栈，按照陷入后厨的方式出栈，执行。
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok2 n1 = ", n2)

	res := n1 + n2
	fmt.Println("ok3 res = ", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res = ", res)
}
