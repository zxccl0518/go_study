package main

import (
	"fmt"
	"time"
)

func main() {
	// 看看日期和时间相关的函数。
	// 1. 获取当前时间。
	now := time.Now() // now 是一个格式化的时间。 \n", now, now)
	// type : time.Time
	// 2019-01-29 14:40:00.4431273 +0800 CST m=+0.011967600

	// 2.通过now 可以获取到年月日，时分秒。
	fmt.Printf("年 = %v\n", now.Year())
	fmt.Printf("月 = %v\n", now.Month())      //月是英文的
	fmt.Printf("月 = %v\n", int(now.Month())) //月是中文的 强转一下就好
	fmt.Printf("日 = %v\n", now.Day())
	fmt.Printf("时 = %v\n", now.Hour())
	fmt.Printf("分 = %v\n", now.Minute())
	fmt.Printf("秒 = %v\n", now.Second())

	// 格式化日期时间。 这种方式 指示输出时间。
	fmt.Printf("当前的年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// 这种方式，可以将格式化好的时间作为字符串 存储到dataStr变量里面。
	dataStr := fmt.Sprintf("当前的年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dataStr = %v， type = %T \n", dataStr, dataStr)

	// 格式化的第二种方式。 使用time提供的方法。Format() 但是这个方法有个要求：这个2006 01 02 15 04 05 这个时间做为标准时间。数字一定不能改。中间的了解符号可以自己改动。
	fmt.Printf(now.Format("2006/01/02 15:04:05")) //2019/01/29 14:59:17
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02")) //2019-01-29
	fmt.Println()
	fmt.Printf(now.Format("15_04_05")) //14_59_17
	fmt.Println()

	// 休眠 time.Sleep()
	// time.Hour:时 time.Minute:分 time.Second:秒  time.Millisecond:毫秒  time.Microsecond:微秒 Nanosecond: 纳秒
	// i := 1
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	time.Sleep(time.Second)
	// 	if i == 11 {
	// 		break
	// 	}
	// }

	// 获取当前unix时间戳、unixnano时间戳。 (作用是可以获取随机数字。)
	fmt.Printf("unix 时间戳=%v, unixnano时间戳=%v \n", now.Unix(), now.UnixNano())
}
