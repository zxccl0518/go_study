package main

import "fmt"

//一个函数test
func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test n1 = ", n1) // ?输出结果
}

func getSum(a int, b int) int {
	sum := a + b
	fmt.Println("test sum :", sum)
	return sum
}

// 请编写函数，可以计算两个数的和与差，并返回结果，
func getSumAndSub(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func main() {
	n1 := 0
	// 调用test
	test(n1)
	fmt.Println("main n1 = ", n1) // ?输出结果

	n2 := getSum(10, 20)
	fmt.Println("main n2 :", n2)

	// 调用 getSumAndSub（）
	// res1, _ := getSumAndSub(30, 20) // 如果我仅仅想要 第一个返回值，第二个不需要，可以使用"_" 来占位。
	res1, res2 := getSumAndSub(30, 20)
	fmt.Printf("res1 = %d, res2 = %d\n", res1, res2)
}
