package main

import "fmt"

// 递归调用的演示 函数。
func test(n int) {
	if n > 2 {
		n--
		test(n)
	}

	fmt.Println("n = ", n)
}

// 地柜方式的演示函数。
func test2(n int) {
	if n > 2 {
		n-- // 递归 必须向退出递归条件逼近，否则就是无限循环。  当执行一个函数是，就创建一个新的受保护的独立空间(新函数栈)
		test2(n)
	} else {
		fmt.Println("n = ", n)
	}
}

func main() {
	// 看一段代码, 结果是 2、2、3
	// test(4)

	// 再看那一段代码。 结果是2
	test2(4)
}
