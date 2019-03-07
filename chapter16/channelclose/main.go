package main

import "fmt"

func main() {
	// intChan := make(chan int, 3)
	// intChan <- 100
	// intChan <- 200

	// close(intChan) // close
	// // 这是不能够在写入数据到channel
	// intChan <- 1 //panic: send on closed channel
	// fmt.Println("okokok")

	newChanint := make(chan int, 100)
	for i := 0; i < 100; i++ {
		newChanint <- i * 2
	}

	// 这种方式 会读取出来一半的内容，因为管道的长度是 动态该变化的。
	// for i := 1; i <= 200; i++ {
	//
	// }

	// 因为 newChanint管道没有close fatal error: all goroutines are asleep - deadlock!
	// 关闭channel 则会正常遍历数据，便利完后，就会退出遍历。
	// 遍历管道 要用for range的方式。
	close(newChanint)
	for v := range newChanint {
		fmt.Println("newChanint = ", v)
	}

}
