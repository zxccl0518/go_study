package main

import "fmt"

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

func main() {
	//  defer 的用法。
	// Go_defer()

<<<<<<< HEAD
	// Pase_student()

	// 二叉树的创建。
	var root *Tree
	t := Create(root, Item{89})
	root = t
	iarr := []int{1, 89, 44, 98, 54, 24, 96, 34, 74, 69, 96, 4, 0}
	for _, v := range iarr {
		create(root, Item{v})
	}
	PreOrder(root)

=======
	Pase_student()
>>>>>>> 9a1353f059a236638f4dc32dcd5d05dec79cd295
}
