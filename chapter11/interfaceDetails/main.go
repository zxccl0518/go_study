package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
}

func (s Stu) Say() {
	fmt.Println("stu Say()")
}

type Integer int

func (i Integer) Say() {
	fmt.Println("integer say i = ", i)
}

type BInterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("实现BInterface接口， monster结构体， Hello方法。")
}
func (m Monster) Say() {
	fmt.Println("实现AInterface接口， monster结构体， Say方法。 ")
}

func main() {
	var stu Stu

	// 因为Stu结构体实现了 AInterface接口，所以可以将一个结构体变量赋给 一个接口定义的变量。
	// 如果 Stu没有实现Say方法，那么是不能 赋值的。Stu does not implement AInterface (missing Say method)
	// 接口本身 不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量(实例)
	// 只要是 自定义类型，都可以是实现接口。不仅仅是结构体。
	// 一个自定义类型 可以实现多个接口。
	// golang中 接口不允许有任何变量。

	var a AInterface = stu
	a.Say()

	// 基本数据类型 作为自定义类型 实现了接口的演示案例。
	var i Integer = 10
	var b AInterface = i
	b.Say()

	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster

	a2.Say()
	b2.Hello()
	fmt.Println("-----------------------------------------")
	monster.Say()
	monster.Hello()
}
