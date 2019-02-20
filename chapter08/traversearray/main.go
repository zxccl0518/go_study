package main

import "fmt"

func main() {
	// 遍历二维数组。
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// for 循环遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}

	// for range方式来遍历二维数组。
	for i, v := range arr {
		for j, k := range v {
			fmt.Printf("arr[%v][%v] = %v \t", i, j, k)
		}
		fmt.Println()
	}
}
