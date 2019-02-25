package main

import "fmt"

// 声明 接口
type BirdAble interface {
	Flying()
}

// Monkey结构体
type Monkey struct {
	Name string
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树 。。。")
}

type LittleMonkey struct {
	Monkey // 匿名结构体
	// m Monkey // 非 匿名结构体
}

// 让LittleMonkey实现BirdAble接口。
func (lm *LittleMonkey) Flying() {
	fmt.Println(lm.Name, "通过接口学习， 会飞  飞飞飞。。")
}

func main() {
	// 创建一个LittleMonkey 实例
	// monkey := &LittleMonkey{
	// 	Monkey{"孙悟空"}}

	monkey := &LittleMonkey{Monkey{"孙悟空"}}
	monkey.climbing() // 如果定义结构体的时候，继承方式是通过匿名结构体的方式继承的，那么调用方法的时候 直接调用。
	// monkey.m.climbing() // 如果不是匿名结构体的方式继承的，那么调用方法的时候就要指明 那个结构体变量 来调用哪个方法。

	// LittleMonkey 实现小鸟的接口后，会飞。
	monkey.Flying()

	// 总结
	/*
		1. 当A结构体继承了B结构体，那么A结构体就自动继承了B结构体的字段和方法，并且可以直接使用。
		2. 当A结构体 需要扩展功能，同时不希望去 破坏继承关系，则可以去实现某个接口即可，因此我们可以认为：实现接口是对继承机制的补充。
	*/
}
