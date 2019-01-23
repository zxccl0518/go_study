package main

import "fmt"

func main() {
	// do while的 golang实现方式。
	// 不同于while ，因为会至少执行一次。先执行，在判断。

	var i int = 1
	for {
		fmt.Println("hello world")
		i++
		if i > 10 {
			break
		}
	}
}
