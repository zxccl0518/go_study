package main

import "fmt"

func main() {
	// 一道面试题，有2个变量，a 和 b，要求将其进行交换，但是不允许试用中间变量。
	var i = 10
	var j = 20
	fmt.Printf("改动之前， i = %v, j = %v\n", i, j)

	// 我自己的答案, 答法错误，启用了中间量。
	var ptrI = &i
	var ptrJ = &j
	*ptrI = j
	*ptrJ = i
	fmt.Printf("改动之 后， i = %v, j = %v\n", i, j)

	// 老师的答案。
	i = i + j
	j = i - j
	i = i - j
	fmt.Printf("改动之 后， i = %v, j = %v\n", i, j)
}
