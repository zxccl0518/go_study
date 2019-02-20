package main

import "fmt"

// 冒泡排序
func BubbleSort(arr *[5]int) {
	var temp = 0
	// 递增的方式排序
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < (len(*arr) - i - 1); j++ {
			if (*arr)[j] < (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func main() {
	// 定义一个数组
	var arr = [5]int{24, 69, 80, 57, 13}
	fmt.Println("arr = ", arr)
	BubbleSort(&arr)
	fmt.Println("BubbleSort() -- arr = ", arr)
}
