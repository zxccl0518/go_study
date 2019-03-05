package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前电脑的cpu最大数
	num := runtime.NumCPU()
	fmt.Println("当前电脑的 最大cpu个数  = ", num)

	// 设置电脑当前cpu 为全部核都开启。
	runtime.GOMAXPROCS(num)
}
