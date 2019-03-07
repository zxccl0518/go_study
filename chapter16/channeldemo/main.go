package main

import "fmt"

type Cat struct {
}

func main() {
	// 演示一下管道的使用
	// 1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 看看 管道是什么
	fmt.Printf("intChan 的值=%v, intChan本身的地址=%p\n", intChan, &intChan)

	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 看看管道的长度和cap容量. 管道容量不会自增长，多了就爆了
	fmt.Printf("channel len = %v  cap = %v\n", len(intChan), cap(intChan))

	var num2 int
	num2 = <-intChan
	fmt.Println("num2 = ", num2)
	fmt.Printf("channel len = %v  cap = %v\n", len(intChan), cap(intChan))

	//在没有使用 协程的情况下，如过我们的管道数据已经全部取出，再去就会报告deadlock

	num3 := <-intChan
	// num4 := <-intChan
	// fmt.Println("num 3 = ", num3, "  < --- > num4 = ", num4) //没有内容继续取 报错： fatal error: all goroutines are asleep - deadlock!

	// 被断言的变量 是接口类型，或者是非接口类型，但是实现了 类型的接口。
	var value interface{}
	value = num3
	if v, ok := value.(int); ok {
		fmt.Printf("value 类型是%T,  int v = %v\n", v, v)
	}
}
