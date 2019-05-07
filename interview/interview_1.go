package main

import (
	"fmt"
	"strconv"

	"github.com/zxccl0518/go_study/interview/closure"
)

// 测试 defer的调用顺序。
func Go_defer() {
	defer func() {
		fmt.Println("打印qian")
	}()
	defer func() {
		fmt.Println("打印 zhong ")
	}()
	defer func() {
		fmt.Println("打印中 hou")
	}()

	// panic("触发异常")

	// 	打印中 hou
	// 打印 zhong
	// 打印qian
	// panic: 触发异常
}

// 采用 defer+recover的方式捕获异常。
func deferAndRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("系统 抛出异常 err = %v\n", err)
		}
	}()

	n1, n2 := 10, 0
	res := n1 / n2
	fmt.Print("res = ", res)
}

type student struct {
	Name string
	Age  int
}

func Pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "wang", Age: 34},
		{Name: "li", Age: 44},
	}

	// for range 的方式遍历，采用的是副本的方式。
	// 这种方式是错误的。
	for _, stu := range stus {
		fmt.Printf("stu dizhi = %p\n", &stu)
		m[stu.Name] = &stu // 这个操作实际上指向的是同一个指针，最终这个指针的值为遍历的最后一个struct的值拷贝。
	}

	// 下面演示 正确的方式。
	// for i := 0; i < len(stus); i++ {
	// 	m[stus[i].Name] = &stus[i]
	// }

	for k, v := range m {
		fmt.Printf("key = %v, value = %v\n", k, v)
	}
}

// 闭包函数测试。
func closureTest() {
	closureObject := closure.GetClosureInstance()

	// 获取闭包函数。
	c := closureObject.MakeSuffix(".jpg")

	fmt.Printf("原文件名字是= %v, 返回的文件名字是 = %s \n", "今天测试", c("今天测试"))
	fmt.Printf("原文件名字是= %v, 返回的文件名字是 = %s \n", "ceshi1", c("ceshi1"))
	fmt.Printf("原文件名字是= %v, 返回的文件名字是 = %s \n\n", "ceshi2.jpg", c("ceshi2.jpg"))

	mp3C := closureObject.MakeSuffix("mp3")
	fmt.Printf("原文件名字是= %v, 返回的文件名字是 = %s \n", "ceshi 3", c("ceshi 3"))
	fmt.Printf("原文件名字是= %v, 返回的文件名字是 = %s \n", "ceshi 4", mp3C("ceshi 4"))

}

// 遍历字符串,包括中文的字符串。
func stringTravelTest() {
	// str := "hello 北京，hello 上海，hello 沈阳，hello 辽阳。"
	// strSlice := []rune(str)
	// for i := 0; i < len(strSlice); i++ {
	// 	fmt.Printf(" %c", strSlice[i])
	// 	if strSlice[i] == '，' {
	// 		fmt.Println()
	// 	}
	// }

	// strconv.Atoi()测试
	n, err := strconv.Atoi("123123123")
	// n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("字符串 转换 int整型 失败 err = ", err)
	} else {
		fmt.Println("转换成功之后的结果是 = ", n)
	}
}

// 切片的创建 demo测试。
func sliceDemo() {
	var sliceInt = make([]int, 5, 10)
	fmt.Println("sliceInt = ", sliceInt)

	sliceInt[0] = 10
	sliceInt[3] = 101
	// sliceInt[5] = 1011 // 这个时候回出现 越界的问题。因为已经超出了初始化的范围了。后面的递增的要append来增加。
	sliceInt = append(sliceInt, 1011)
	fmt.Println("sliceInt = ", sliceInt)
}

func main() {
	//  defer 的用法。
	// Go_defer()

	// Pase_student()

	// 二叉树的创建。
	// var root *Tree
	// t := Create(root, Item{89})
	// root = t
	// iarr := []int{1, 89, 44, 98, 54, 24, 96, 34, 74, 69, 96, 4, 0}
	// for _, v := range iarr {
	// 	create(root, Item{v})
	// }
	// PreOrder(root)

	// Pase_student()

	// closureTest()

	// 字符串的遍历，包括中文的字符串。
	// stringTravelTest()

	// deferAndRecover()

	// 切片的测试代码。
	sliceDemo()
}
