package main

import "fmt"

func InserSort(arr *[5]int) {
	// 完成第一次，给第二个元素找到合适的位置并插入。
	// inserVal := arr[1]
	// insertIndex := 1 - 1

	// for insertIndex >= 0 && arr[insertIndex] < inserVal {
	// 	arr[insertIndex+1] = arr[insertIndex] // 数据后移
	// 	insertIndex--
	// }
	// if insertIndex+1 != 1 {
	// 	arr[insertIndex+1] = inserVal
	// }

	// // 完成第二次 给第三个元素找到合适的位置。
	// inserVal = arr[2]
	// insertIndex = 2 - 1

	// for insertIndex >= 0 && arr[insertIndex] < inserVal {
	// 	arr[insertIndex+1] = arr[insertIndex] // 数据后移
	// 	insertIndex--
	// }
	// if insertIndex+1 != 2 {
	// 	arr[insertIndex+1] = inserVal
	// }

	// 以上以此类推。数组n个元素，那么需要循环n-1次
	for i := 1; i < len(arr); i++ {
		insertValue := arr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && arr[insertIndex] < insertValue {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}

		if insertIndex+1 != i {
			arr[insertIndex+1] = insertValue
		}
	}
}

func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	fmt.Println("charu zhiqian arr  = ", arr)
	InserSort(&arr)
	fmt.Println("charu zhihou arr  = ", arr)
}
