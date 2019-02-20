package main

import "fmt"

func main() {
	// 二维数组的演示案例。
	/*
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/

	//定义 / 声明二维数组。
	var arr [4][6]int
	// 赋值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	fmt.Println("arr = ", arr) //arr =  [[0 0 0 0 0 0] [0 0 1 0 0 0] [0 2 0 3 0 0] [0 0 0 0 0 0]]

	// 遍历二位数组，按照要求输出图形。
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}
