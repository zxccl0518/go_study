package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	// 在实行test03前，先获取到当前的unix时间戳
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("test03 消耗时间： %v 秒\n", end-start)
}
