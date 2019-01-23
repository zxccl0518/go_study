package main

import "fmt"

func main() {
	// 要求：可以从控制台接受用户信息，【姓名，年龄，薪水，是否通过考试。】
	var name string
	var age byte
	var sal float32
	var isPass bool
	// 方式 1. fmt.Scanln
	// fmt.Println("请输入姓名:")
	// // 当程序执行到 fmt.Scanln(&name, 程序会停止在这里，等待用户输入，并回车。
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水:")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入 是否通过考试:")
	// fmt.Scanln(&isPass)
	// fmt.Printf("名字是 %v\n年龄是:%v\n薪水是：%v\n是否通过考试：%v", name, age, sal, isPass)

	// 方式2 Scanf()
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，使用空格分隔开。")
	fmt.Scanf("%s %d %f %v", &name, &age, &sal, &isPass)
	fmt.Printf("名字是 %v\n年龄是:%v\n薪水是：%v\n是否通过考试：%v", name, age, sal, isPass)
}
