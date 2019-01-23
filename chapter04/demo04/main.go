package main

import "fmt"

func main() {
	// 演示逻辑运算符的使用。   -> &&
	var age = 40

	if age > 30 && age < 50 {
		fmt.Println("&& ok1")
	}

	if age > 30 && age < 40 {
		fmt.Println("&& ok 2")
	}

	// 演示逻辑运算符的使用。   -> ||
	if age > 30 || age < 50 {
		fmt.Println("|| ok1")
	}

	if age > 30 || age < 40 {
		fmt.Println("|| ok 2")
	}

	// 演示逻辑运算符的使用。   -> !
	if !(age > 30) {
		fmt.Println("!  ok ")
	}

}
