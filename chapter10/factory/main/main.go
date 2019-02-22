package main

import (
	"fmt"

	"github.com/zxccl0518/Go_dev/chapter10/factory/main/model"
)

func main() {
	// 创建一个Student实例
	// var stu = model.Student{ // 这种写法只有model包中的机构提首字母是大写的情况下才可以使用。
	// 	Name:  "tom",
	// 	Score: 78.9,
	// }
	// fmt.Println(stu)

	var stu1 = model.NewStudent("lilei", 99.1)
	var stu2 = model.NewStudent("lilei", 99.1)
	fmt.Println("采用工厂模式的方式创建的结构体指针 stu = ", stu1) // 等价于(*stu)
	// fmt.Println("采用工厂模式的方式创建的结构体指针 stu.name = ", stu.Name, " stu.Score = ", stu.Score)// 仅仅结构体中 定义的字段的首字母是大写的才可以 直接调用。

	// 现在是字段的首字母 为小写的时候 调用方式演示。
	fmt.Println("现在是 结构体中 字段的首字母是小写的情况，【地址传递】调用字段的方式 stu.name = ", stu1.GetName(), ", stu.score = ", stu1.GetScore())
	fmt.Println("现在是 结构体中 字段的首字母是小写的情况，【值  传递】调用字段的方式 stu.name = ", (*stu2).GetName1(), ", stu.score = ", (*stu2).GetScore1())
	fmt.Println("【地址传递】调用字段的方式  stu1 = ", stu1)  //【地址传递】调用字段的方式  stu1 =  &{haha 10.1}
	fmt.Println("【值  传递】调用字段的方式  stu2 = ", stu2) //【值  传递】调用字段的方式  stu2 =  &{lilei 99.1} 值传递方式 即使在方法里面改动，外面的结构体实例也不会受到改动
}
