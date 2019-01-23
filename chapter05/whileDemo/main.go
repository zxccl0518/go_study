package main

import "fmt"

func main() {

	// 使用while的方式输出10局 hello,world
	// 循环变量初始化。
	var i int = 1
	for {
		if i > 10 {
			break
		} else {
			fmt.Println("hello world")
			i++
		}
	}
}
