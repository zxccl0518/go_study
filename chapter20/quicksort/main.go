package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 快速排序。
// left表示数组左边的小标
// right表示数组右边的下表
// array表示要排序的数组s'd
func quick(left int, right int, array *[800]int) {
	l := left
	r := right
	pivot := array[(left+right)/2]
	// temp := 0

	// for 循环的目标是将比pivot肖的书放在左边。
	// 比pivot大数放到右边。
	for l < r {
		// 从pivot的左侧找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}

		// 从pivot的右侧找到小于等于pivot的值。
		for array[r] > pivot {
			r--
		}

		// 表明本次分解任务完成，break
		if l >= r {
			break
		}

		// 交换
		array[l], array[r] = array[r], array[l]

		// 优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}

		if left < r {
			quick(left, r, array)
		}

		if right > l {
			quick(l, right, array)
		}
	}
}

func main() {
	// arr := [6]int{-9, -78, 0, 23, -567, 70}
	var arr [800]int
	for i := 0; i < 800; i++ {
		arr[i] = rand.Intn(900000)
	}

	start := time.Now().Unix()
	// fmt.Println("arr origal = ", arr)
	quick(0, len(arr)-1, &arr)
	// fmt.Println("arr quicksort result = ", arr)
	end := time.Now().Unix()

	fmt.Printf("start = %d, end = %d, total = %d \n", start, end, start-end)
}
