package main

import (
	"fmt"
)

// 计算 是不是素数
func PrimeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	flag := true
	for {
		// time.Sleep(time.Millisecond * 10)

		v, ok := <-intChan
		if !ok {
			break
		}

		flag = true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag == true {
			primeChan <- v
		}
	}

	exitChan <- true
	fmt.Println("协程 读取不到 通道中的内容， 协程退出。")
}

// 将1-8000 循环写入通道内
func PutNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}

	close(intChan)

	// for i := 1; i <= 4; i++ {
	// 	go PrimeNum(intChan, primeChan, exitChan)
	// }
}

//
func main() {
	// 求出 1-8000 之中所有的素数
	goroutineNums := 4
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, goroutineNums)

	go PutNum(intChan, primeChan, exitChan)

	for i := 1; i <= goroutineNums; i++ {
		go PrimeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 1; i <= goroutineNums; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	// 遍历primeChan 把素数取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数 res = ", res)
	}

	fmt.Println("main 主程序退出。 ---")
}
