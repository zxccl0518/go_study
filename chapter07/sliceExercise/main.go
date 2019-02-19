package main

import "fmt"

// 传递的参数是一个引用， 也就是说地址。
func test(slice []int) {
	// slice的改变会影响原来的实参值
	slice[0] = 100
}

func main() {
	var slice []int
	var arr = [5]int{1, 2, 3, 4, 5}
	slice = arr[:]
	slice[0] = 10
	var slice2 = slice
	fmt.Println("slice = ", slice)   //slice =  [10 2 3 4 5]
	fmt.Println("slice2 = ", slice2) //slice2 =  [10 2 3 4 5]
	fmt.Println("arr = ", arr)       //arr =  [10 2 3 4 5]

	// 经过test函数，slice的第一个值 被改编为 100
	test(slice)
	fmt.Println("test() -- slice = ", slice) //test() -- slice =  [100 2 3 4 5]
}
