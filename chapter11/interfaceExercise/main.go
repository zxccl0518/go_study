package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
}

func (s *Stu) Say() {
	fmt.Println("say()")
}

func main() {
	var s = Stu{}
	//报错：Stu does not implement AInterface (Say method has pointer receiver) 因为实现Say的时候 是结构体的指针 形式 实现的。所以 给接口变量赋值的时候 只能赋值指针形式的。
	// var i AInterface = s

	//	正确。
	var i AInterface = &s

	i.Say()
}
