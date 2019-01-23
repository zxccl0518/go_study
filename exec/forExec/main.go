package main

import "fmt"

func main() {
	// 打印1~100之间所有是9的倍数的整数的个数及总和

	// 分析思路。
	// 1.使用for 循环对max进行遍历。
	// 2.当一个数%9 == 0 就是9的倍数
	var max = 100
	var sum int
	for i := 1; i <= max; i++ {
		if i%9 == 0 {
			fmt.Printf("%d ", i)
			sum += i
		}
	}

	fmt.Println("总和 sum = ", sum)
	fmt.Println("------------------------------------------------")

	//完成下面的表示输出，6是可变的。
	var n = 60
	for i := 0; i <= n; i++ {
		fmt.Printf("%v + %v = %v \n", i, n-i, n)
	}
}
