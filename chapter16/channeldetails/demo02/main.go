package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用select可以解决从管道取数据的阻塞问题。

	// 1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2. 定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道的时候，如果不关闭会阻塞 而导致 deadLock
	// 问题，在实际开发中，可能我们不好确定什么时候 应该关闭管道
	// lable:
	for {
		select {
		case v := <-intChan:
			time.Sleep(time.Microsecond * 100)
			fmt.Println("从 intChan 通道里面读取出来的shuju = ", v)
		case v := <-stringChan:
			time.Sleep(time.Microsecond * 100)
			fmt.Println("从 stringChan 通道里面读取出来的shuju = ", v)
		default:
			fmt.Println(" 任何一个 通道中 都无法读取数据出来了。 不玩了 拜拜了您内。")
			// break lable
			return
		}
	}
}
