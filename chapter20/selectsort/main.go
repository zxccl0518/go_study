package main

import "fmt"

// 选择排序。
func SelectSort(arr *[5]int) {
	var maxValue int
	var maxIndex int
	for i := 0; i < len(arr)-1; i++ {
		maxIndex = i
		maxValue = arr[i]

		for j := i + 1; j < len(arr); j++ {
			if maxValue < arr[j] {
				maxValue = arr[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
}

func main() {
	arr := &[5]int{15, 2, 4, 55, 99}

	fmt.Println("排序之前 ：", arr)
	SelectSort(arr)
	fmt.Println("排序之后 ：", arr)
}
