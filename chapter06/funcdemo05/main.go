package main

import "fmt"

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

// 自定义函数类型 案例
type anothorFuncName func(int, int) int // 这事 anothorFuncName 就是 func(int, int) int

// 函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用。
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	// func myFun(funvar anothorFuncName, num1 int, num2 int) int { 等价于上面一行的写法。因为有了自定义函数。
	return funvar(num1, num2)
}

//支持对函数返回值命名。
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return // 直接return 就可以了。
}

// 案例演示，编写一个函数sum，可以求出 1 到多个int的和。
func sumArgs(n1 int, args ...int) sum int {
	// func sumArgs(n1 int, args ...int) int {
	sum := n1
	// 迭代器的方式遍历
	// for index, val := range args {
	// 	fmt.Printf("index = %v, agrs = %v \n", index, val)
	// 	sum += val
	// }

	// 正常的for 循环遍历。
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	// return sum
	return
}

func main() {
	// 看案例。一个函数作为形参 传入另一个函数。
	res := myFun(getSum, 50, 40)
	fmt.Println("res = ", res)

	// 自定义 变量
	// 给int取了别名。 但是在go中，myInt 和 int 虽然都是int类型，但是go认为myInt 和int是两个类型。但是可以通过强制类型转换 使用。
	type myInt int
	var a1 myInt = 100
	var a2 = 100
	fmt.Printf("a 的类型：%T, a2的类型：%T \n", a1, a2)
	// a2 = a1, 错误的写法，因为类型不同。
	a2 = int(a1) // 这个写法是正确的。

	// 自定义函数类型 案例
	// type anothorFuncName func(int, int) int // 这事myFun就是 func(int, int) int
	res2 := myFun(getSum, 2, 3)
	fmt.Println("res2 = ", res2)

	// 对返回值进行命名。案例演示。
	a, b := getSumAndSub(1, 2)
	fmt.Printf("a = %d, b = %d\n", a, b)

	// 可变参数的 案例
	sumValue := sumArgs(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("可变参数的案例， 求和结果 sumValue:", sumValue)
}
