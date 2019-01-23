package main

import "fmt"

func main() {
	// 求出两个数的最大值
	var n1 = 10
	var n2 = 40
	var max int
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("2个数的 max  = ", max)

	// 求出3个值的最大值。
	// 思路：先求出2个数的最大值，再跟第三个数比较。
	var n3 = 45
	if n3 > max {
		max = n3
	}
	fmt.Println("3个数的 max  = ", max)
}
