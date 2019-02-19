package main

import "fmt"

func main() {
	// 使用常规的for循环遍历切片。
	var arr = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	// slice := arr[1:len(arr)]
	// slice := arr[0:len(arr)]  等价于 slice = arr[:] 就是引用数组的全部元素
	fmt.Println("=========================================================")
	sliceTmp1 := arr[:len(arr)]
	fmt.Println("sliceTmp1 = ", sliceTmp1)
	sliceTmp2 := arr[:]
	fmt.Println("sliceTmp2 = ", sliceTmp2)
	fmt.Println("=========================================================")

	for i := 0; i < len(slice); i++ {
		fmt.Printf("for -> slice[%v] = %v \n", i, slice[i])
	}

	// for range方式遍历切片。
	for i, v := range slice {
		fmt.Printf("for range -> slice[%v] = %v \n", i, v)
	}

	// 切片可以继续切片。
	slice2 := slice[1:2]
	fmt.Println("slice2 = ", slice2)

	// 将slice2 的元素改变了。那么slice  和 arr 都将会变化，因为指向的是同一个地址空间。
	slice2[0] = 100
	fmt.Println("slice = ", slice)
	fmt.Println("arr = ", arr)

	// append内饰函数，可以对切片进行动态追加。
	fmt.Printf("slice = %v, slice len = %v ,cap = %v \n", slice, len(slice), cap(slice))
	slice = append(slice, 777, 888, 999, 1000, 1001, 1002)
	fmt.Printf("slice = %v, slice len = %v ,cap = %v \n", slice, len(slice), cap(slice))

	// append 只能对 切片进行追加扩容操作。 append的第一个参数可以自己本身的切片，也可以是其他切片。
	slice3 := append(slice, 1, 2, 3)
	fmt.Printf("slice3 = %v, slice3 len = %v ,cap = %v slice3的地址 = %p \n", slice3, len(slice3), cap(slice3), &slice3)
	slice3 = append(slice3, slice3...) // 第二个参数，可以是个切片 但是一定要 "..." 而且只能有一个 "切片..."
	fmt.Printf("slice3 = %v, slice3 len = %v ,cap = %v slice4的地址 = %p \n", slice3, len(slice3), cap(slice3), &slice3)

	fmt.Println("------------------------------------------------------------")
	// 切片的拷贝操作
	// 切片使用copy内置函数完成拷贝，举例说明。
	var slice4 = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10, 20)
	copy(slice5, slice4)
	fmt.Println("slice4 = ", slice4)
	fmt.Println("slice5 = ", slice5)
	slice4[1] = 100 // 这里改变的话，对slice5也不会有影响，因为copy是 对2个独立的空间进行拷贝。
	fmt.Println("slice4 = ", slice4)
	fmt.Println("slice5 = ", slice5)

}
