package main

import "fmt"

func main() {
	// 赋值运算符的演示。  有一个特点，执行顺序从右向左。
	// var a int
	// a = 10 // 基本赋值

	// 有两个变量，a 和 b，要求将其进行交换，最终打印结果。
	// a = 9， b= 2
	a := 9
	b := 2
	fmt.Printf("交换前的情况是 a= %v, b = %v \n", a, b)
	t := a
	a = b
	b = t
	fmt.Printf("交换后的情况是 a= %v, b = %v\n", a, b)

	// 复合赋值的操作。
	a += 17 // 等价于a = a + 17

	// 2) 赋值运算符的左边，只能事变量，右边可以是变量，表达式，常量值。
	// 表达式：任何有值都可以看作是表达式。
	var d int

	d = a
	d = test() + 90 // 右边是表达式。
	// d = 8 + 2*3
	fmt.Println("d = ", d)
}

func test() int {
	return 10
}
