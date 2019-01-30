package main

import "fmt"

func main() {
	// var heros [3]string = [...]string{"宋江", "吴用", "卢俊义"}s
	heros := [...]string{"宋江", "吴用", "卢俊义"}
	fmt.Println("heros = ", heros)

	// 演示 for-range 遍历数组。
	for i, v := range heros {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}

	// 这么写的话，就是不关心下标识多少，只关心v 也就是关心数组的元素的值。
	for _, v := range heros {
		fmt.Printf(" v = %v\n", v)
	}
}
