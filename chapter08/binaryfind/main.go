package main

import "fmt"

/*
二分查找。
*/
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findValue int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到了。")
		return
	}

	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findValue {
		rightIndex = middle - 1
		BinaryFind(arr, leftIndex, rightIndex, findValue)
	} else if (*arr)[middle] < findValue {
		leftIndex = middle + 1
		BinaryFind(arr, leftIndex, rightIndex, findValue)
	} else {
		// 找到了。
		fmt.Printf("找到了， 下标为:%v\n", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}

	BinaryFind(&arr, 0, len(arr)-1, 1234)
}
