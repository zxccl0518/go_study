package main

import "fmt"

func fbn(n int) []uint64 {
	var slice = make([]uint64, 1)
	if n < 0 {
		slice[0] = 0
		return slice
	}

	if n == 0 {
		slice[0] = 1
	} else if n == 1 {
		slice[0] = 1
		slice = append(slice, 1)
	} else {
		slice[0] = 1
		slice = append(slice, 1)
		for i := 2; i < n; i++ {
			slice = append(slice, (slice[i-2] + slice[i-1]))
		}
	}

	return slice
}

func main() {
	/*
		1)可以接受一个n int
		2)能够将斐波那契额的数列放到切片中，
		3)提示 斐波那契的数列形式:
		arr[0] =1 ,arr[1] =1,arr[2] =2,arr[3] =3,arr[4] =5

		思路：
		1)声明一个函数fbn(n int) ([]uint64)
		2)变成fbn(n int)进行for循环来存放斐波那契额的数列 0=>1, 1=>1 下标为 0 和 1 的有点特别，其他符合前2个数之和为后一个数的规律
	*/
	slice := fbn(1)
	fmt.Println("fbn() slice = ", slice)
}
