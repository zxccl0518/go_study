package main

import "fmt"

func main() {
	// 演示切片的基本使用， 切片是数组的一个引用。
	// 数组
	var intArr = [...]int{1, 2, 3, 4, 5}
	// 声明、定义一个切片。
	// 1. slice 就是切片名
	// 2. intArr[1:3 表示slice 引用到intArr 这个数组。
	// 3. 引用intArr数组的起始下标 为1，终止下标为3，但是不包括3 。
	slice := intArr[1:3]
	fmt.Println("intArr = ", intArr)          //[1 2 3 4 5]
	fmt.Println("slice 元素 = ", slice)         //[2 3]
	fmt.Println("slice 长度len = ", len(slice)) //2
	fmt.Println("slice 的容量 = ", cap(slice))   //4, 切片容量是可以动态变化的。

	fmt.Printf("intArr[1]的地址 = %p, intArr[2]的地址是:%p \n", &intArr[1], &intArr[2])
	fmt.Printf("slice[0]的地址 = %p, slice[0] = %v, slice[1]的地址是:%p \n", &slice[0], slice[0], &slice[1])
	// intArr[1]的地址 与 slice[0] 地址是 相同的。

	// slice 从底层的层面来说 就是一个strut的结构体。

	// 由于 切片是数组的引用，一旦切片改变了某个元素的值，那么数组 和 切片的元素值都会变化。
	slice[1] = 100
	fmt.Println("\nintArr = ", intArr)
	fmt.Println("slice = ", intArr)

	// 总结：
	// 1.slice的确是一个引用类型。
	// 2.slice从底层来说，其实就是一个数据结构struct结构体。
	// type slice struct{
	//  ptr *[num]int		一个地址的指针
	//  len int             当前的长度
	//  cat int             当前的容量。
	// }

}
