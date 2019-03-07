package main

import (
	"fmt"
	"sync"
	"time"
)

// 要求计算 1-200各个数的阶乘，并把各个数的阶乘放入到map中。
// 最后显示出来，要求使用goroutine完成。

// 思路
// 1.编写一个函数，来计算各个数的阶乘，并放入map中。
// 2.我们启动的协程多个，统一的将阶乘的结果放入到map中。
// 3.map应该做成全局的。

var (
	m = make(map[int]int, 10)
	// 声明一个 全局的互斥锁
	// lock是一个全局的互斥锁，
	// sync是一个包，是同步的意思，全称是synchornized 同步
	// Mutex 是互斥的意思。
	lock sync.Mutex
)

// 计算n的阶乘
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	// 加锁
	lock.Lock()
	m[n] = res
	// 解锁
	lock.Unlock()
}

func main() {
	var i int
	for i = 0; i < 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)

	// 主线程休眠10秒，为的是让协程完成任务
	// 如果没有这句话的话，主线程很快就退出，m中还没有结果呢。
	lock.Lock()
	for i = 0; i < len(m); i++ {
		fmt.Printf("m[%v] = %v \n", i, m[i])
	}
	lock.Unlock()
}
