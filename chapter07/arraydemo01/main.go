package main

import "fmt"

func main() {
	/*
		一个养鸡场有6只鸡，他们的体重分别是3kg， 5kg，1kg，3.4kg，2kg，50kg
		请问这六只鸡的总体中是多少？ 平均体重是多少？请你编写一个程序. => 数组。
	*/

	// 思路分析，定义六个变量，分别表示六只鸡，然后求出和，然后求出平均值。
	// hen1 := 3.0
	// hen2 := 5.0
	// hen3 := 1.0
	// hen4 := 3.4
	// hen5 := 2.0
	// hen6 := 50.0
	// totalWeight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	// avgWeight := totalWeight / 6

	// fmt.Printf("totalWeight = %v, avgWeight = %v \n", totalWeight, avgWeight)   //totalWeight = 64.4, avgWeight = 10.733333333333334
	// fmt.Printf("totalWeight = %v, avgWeight = %.2f \n", totalWeight, avgWeight) //totalWeight = 64.4, avgWeight = 10.73,  .2f 可以保留小数点后2位。但是仅仅是以小数点后2位的形式输出，变量本质还是原来的位数。

	// avgWeight1 := fmt.Sprintf("%.2f", totalWeight/6)                           // 可以直接将得到的结果 进行2位小数保存。
	// fmt.Printf("totalWeight = %v, avgWeight = %v \n", totalWeight, avgWeight1) //totalWeight = 64.4, avgWeight = 10.733333333333334

	// 特别注意：数组是值类型，而非其他语言的 引用类型。需要保证同意数组的变量类型相同。

	// 使用数组的方式 来解决问题。
	// 1.定义一个数组，
	var hens [6]float64
	// 2.给数组的每个元素赋值。
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	// 3.遍历数组，求出总体重
	totalWeight2 := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight2 += hens[i]
	}
	// avgWeight2 := totalWeight2 / float64(len(hens))
	// fmt.Printf("totalWeight = %v, avgWeight = %v \n", totalWeight2, avgWeight2) //totalWeight = 64.4, avgWeight = 10.733333333333334

	avgWeight2 := fmt.Sprintf("%.2f", totalWeight2/float64(len(hens)))
	fmt.Printf("totalWeight = %v, avgWeight = %v \n", totalWeight2, avgWeight2) //totalWeight = 64.4, avgWeight = 10.73
}
