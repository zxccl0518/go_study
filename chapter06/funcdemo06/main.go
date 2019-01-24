package main

import "fmt"

// 累加器
// func AddUpper() func(x int) int {
// 	var n = 10
// 	fmt.Println("==1==> n = ", n)
// 	return func(x int) int {
// 		fmt.Println("==2==> n = ", n)
// 		n = n + x
// 		return n
// 	}
// }

func AddUpper1() func(x int) int {
	var n = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += "A"
		fmt.Println("str = ", str)
		return n
	}
}

// 全局匿名函数 定义。
var (
	Func1 = func(n1 int, n2 int) (sum int) {
		sum = n1 * n2
		return
	}
)

func main() {
	// 如果是下面哪种调用方式，则不会起到 累加器的作用，因为每次都初始化AddUpper这个方法，里面的n每次都会初始化。
	// fmt.Println(AddUpper()(0))
	// fmt.Println(AddUpper()(1))
	// fmt.Println(AddUpper()(2))
	// fmt.Println(AddUpper()(3))

	// 闭包： 闭包就是一个函数 和与其相关的引用环境 组合的一个整体(实体)
	// f1 := AddUpper() // 通过这种方式，仅仅初始化一次函数，然后将函数赋给一个变量，再调用这个变量。就可以实现叠加器的功能。

	// fmt.Println(f(1))
	// fmt.Println(f(2))
	// fmt.Println(f(4))

	//闭包的改动版本。
	f2 := AddUpper1()
	fmt.Println("闭包测试的第二个版本  :", f2(1))
	fmt.Println("闭包测试的第二个版本  :", f2(2))
	fmt.Println("闭包测试的第二个版本  :", f2(3))

	// 匿名函数的演示案例。

	/*
		// 定义匿名函数的同时，就调用了这个匿名函数。 一次性使用。
		res1 := func(n1 int, n2 int) int {
			return n1 + n2
		}(1, 2)
		fmt.Println("res1 : ", res1)

		// 将匿名函数 func(n1 int, n2 int) int 赋给a变量
		// 则a的数据类型就是函数类型，此时
		a := func(n1, n2 int) int {
			return n1 - n2
		}
		subValue1 := a(10, 20)
		fmt.Println("subValue1 : ", subValue1)
		subValue2 := a(1, 2)
		fmt.Println("subValue2 : ", subValue2)

		//全局匿名函数 测试
		result := Func1(2, 3)
		fmt.Println("result = ", result)
	*/
}
