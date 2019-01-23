package main

import "fmt"

func main() {
	var num int = 9
	fmt.Printf("num address: %v \n", &num)

	var ptr *int
	ptr = &num
	*ptr = 100 // 这里修改时，会影响num的值的变化
	fmt.Println("new nun is ", *ptr)
}
