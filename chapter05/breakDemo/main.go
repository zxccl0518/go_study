package main

import (
	"fmt"
)

func main() {
	// 我们为了生成一个随机数，还需要个rand设置一个种子。
	// time.Now().unix() : 返回一个从1970.01.01 的0时0分0秒到现在的秒数。
	// rand.Seed(time.Now().Unix())
	// 如何随机的生成 1-100整数。
	// n := rand.Intn(100) + 1 // [0,100)
	// fmt.Println("随机数是 ：", n)

	// 随机生成一个1-100 的随机数，直到生成99这个数为止，看看用过了几次。
	// 分析思路：
	// 编写一个无限循环的控制，然后不停的生成随机数，当生成了99的时候，就退出这个无限循环。 break

	// rand.Seed(time.Now().Unix())			//秒级
	// rand.Seed(time.Now().UnixNano()) // 纳秒
	// var value int
	// var counts int
	// for {
	// 	value = rand.Intn(100) + 1
	// 	if value == 99 {
	// 		fmt.Printf("这个数是 ：%d ,这次随机一共用了%d 次", value, counts)
	// 		break
	// 	}
	// 	counts++
	// }

	// break 可以通过制定标签 来决定跳出那个循环。
lable2:
	for i := 1; i < 4; i++ {
	lable1:
		for j := 0; j < 10; j++ {
			if j == 7 {
				break lable2
			} else if j == 5 {
				fmt.Println("测试 break lable2 还会不会打印这条log")
				break lable1
			}
		}
	}
	fmt.Println("break lable2")
}
