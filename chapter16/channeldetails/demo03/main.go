package main

import (
	"fmt"
	"time"
)

// 函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello . world")
	}
}

// 函数
func test() {
	// 这里我们可以使用defer + recover
	defer func() {
		// 捕获 test跑出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生 错误. err = ", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"
}
func main() {
	go sayHello()

	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("main() ok . i = ", i)
	}
}
