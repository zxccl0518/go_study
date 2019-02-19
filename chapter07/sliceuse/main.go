package main

import "fmt"

func main() {
	// 演示切片的使用。 make
	// var slice []float64
	// fmt.Println("slice = ", slice) //slice =  []。 slice切片为空

	// 对于切片而言，必须make之后才可以使用。******
	var slice []float64 = make([]float64, 5, 10)
	fmt.Println("slice = ", slice)          //slice =  [0 0 0 0 0]
	fmt.Println("slice len = ", len(slice)) //slice len =  5
	fmt.Println("slice cap = ", cap(slice)) //slice cap =  10

	//1) 通过make的方式创建切片可以指定切片的大小和容量。
	//2) 如果没有给切片的各个元素赋值，那么就会使用默认值【int,float=> 0  string = ""  bool = false】
	//3) 通过make的方式创建的切片对应的数组是由make的底层来维护的。对外不可见。即 只能通过sliece去访问各个元素。

	// slice的第三种使用方式。:定义一个切片，直接就指定具体数组，使用原理类似make的方式。
	var strslice = []string{"tom", "jack", "tony"}
	fmt.Println("strslice = ", strslice)          //strslice =  [tom jack tony]
	fmt.Println("strslice len = ", len(strslice)) //strslice len =  3
	fmt.Println("strslice cap = ", cap(strslice)) //strslice cap =  3
	fmt.Printf("strSlice type = %T \n", strslice)

	// 一道面试题
	//  方式1：var intArr = [5]int {1,2,3,4,5} slice := intArr[1:3]
	//  方式2：var intArr = []int {1,2,3,4,5}
	// 作答：1) 方式1 是直接引用数组，这个数组是事先存在的，程序员是可见的。
	//       2) 方式2 是通过make来创建切片的，make也会创建一个数组，是由切片在底层进行维护，程序员是看不见的
	var intArr1 = [5]int{1, 2, 3, 4, 5}
	slice1 := intArr1[1:3]
	fmt.Println("slice1 = ", slice1) //slice1 =  [2 3]

	var slice2 = []int{1, 2, 3, 4, 5} //slice2 =  [1 2 3 4 5]
	fmt.Println("slice2 = ", slice2)

}
