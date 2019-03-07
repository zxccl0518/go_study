package main

import (
	"fmt"
	"time"
)

type Cat struct {
	Name string
	Age  int
}

// write data
func WriteData(inChan chan int) {
	for i := 1; i <= 50; i++ {
		fmt.Println("writeData i = ", i)
		inChan <- i
	}

	fmt.Println("关闭 inChan 通道")
	close(inChan)
}

// read data
func ReadData(inChan chan int, outChan chan bool) {
	for {
		v, ok := <-inChan
		if !ok {
			break
		}
		fmt.Println("--1--> Readdata v = ", v)
		time.Sleep(1000 * time.Nanosecond)
	}

	outChan <- true
	fmt.Println("关闭 outChan 通道")
	close(outChan)
}

//
func ReadData1(inChan chan int, outChan chan bool) {
	for {
		v, ok := <-inChan
		if !ok {
			break
		}
		fmt.Println("--2--> Readdata v = ", v)
		time.Sleep(1000 * time.Nanosecond)
	}
}

func main() {
	var inChan = make(chan int, 10)
	var outChan = make(chan bool, 1)
	go WriteData(inChan)
	// go ReadData(inChan, outChan)	// 如果程序中 对于管道 只写入 而没有读取操作，就会出现死锁的问题，哪怕是有另一个地方会读 读的很慢也无所谓。只要有读的地方 就会不出现死锁。
	// go ReadData1(inChan, outChan)

	// 因为在协程中有 数据在写入,所以阿合力的管道读取的操作会阻塞住。直到 读到内容位置。
	for {
		if value, ok := <-outChan; ok {
			fmt.Println("读取到值了。 value = ", value)
			break
		} else {
			fmt.Println("meiyou 读取到值了。  ")
		}
	}

	// channel := make(chan interface{}, 10)
	// cat1 := Cat{"miao 1", 10}
	// cat2 := Cat{"miao 22", 20}
	// channel <- cat1
	// channel <- cat2
	// channel <- 100
	// channel <- "hahah"
	// <-channel
	// <-channel
	// <-channel
	// cat11 := <-channel
	// newCat, ok := cat11.(Cat)
	// if ok {
	// 	fmt.Println("cat 11 name = ", newCat.Name)
	// } else {
	// 	fmt.Println("cat11 不是 Cat结构体类型的变量。不能读取到Name 字段。")
	// }

	// intChan := make(chan int, 100)
	// go func(intChan chan int) {
	// 	// for {
	// 	time.Sleep(5 * time.Second)
	// 	// }
	// 	intChan <- 1
	// 	// close(intChan)	// 如果不加入close 关闭管道，则会在遍历管道的时候报错 deadlock死锁
	// }(intChan)

	// for {
	// 	v, ok := <-intChan
	// 	if ok {
	// 		fmt.Println("读取到了 === v = ", v)
	// 	} else {
	// 		fmt.Println("没有到了 ==xxx==")
	// 		break
	// 	}
	// }
}
