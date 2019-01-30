package main

import "fmt"

func main() {
	var intArr [3]int
	// 当我们定义完数组后，其实数组的各个元素已经有默认值了。
	for i := 0; i < len(intArr); i++ {
		fmt.Printf("intArr[%d] = %v\n", i, intArr[i])
	}
	// 默认值是 [0, 0, 0]

	// 数组的地址 就是其首个元素的地址。 数组元素的地址是 连续的。
	fmt.Printf("intArr的地址=%p, intArr[0]的地址=%p\n", &intArr, &intArr[0])
	// 那么数组第二个元素的地址是多少呢? 是这个元素占的字节数。
	fmt.Printf("intArr的地址=%p, intArr[1]的地址=%p\n", &intArr, &intArr[1])
	// 确实第一个 第二个地址相差8 ，也就是8个字节。。 如果当初 数组是 var intArr [3]int32，那么每个元素之间的地址相差是4个字节。
	// int 8个字节。  int32 4个字节。

	// 几种初始化数组的方式
	// var numArray01 [3]int = [3]int{1, 2, 3}
	var numArray01 = [3]int{1, 2, 3}
	fmt.Println("numArray01 = ", numArray01)

	numArray02 := [3]int{4, 5, 6}
	fmt.Println("numArray02 = ", numArray02)

	var numArray03 = [...]int{7, 8, 9} // [...] 是固定写法，就是说程序自己去判断自己有多少个元素。
	fmt.Println("numArray03 = ", numArray03)

	var names = [3]string{1: "apple", 0: "banana", 2: "orange"}
	fmt.Println("names = ", names)
}
