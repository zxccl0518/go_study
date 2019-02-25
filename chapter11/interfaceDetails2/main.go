package main

import "fmt"

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Stu struct {
}

func (s Stu) test01() {
	fmt.Println("test01 ---test01")
}

func (s Stu) test02() {
	fmt.Println("test02 ---test02")
}
func (s Stu) test03() {
	fmt.Println("test02 ---test02")
}

// 空接口 展示
type T interface{}

func main() {
	var s = Stu{}
	var a AInterface = s // a接口 集成B C接口，全部实现B C接口的所有方法才可以 使用a 结构 对应的实例。

	// 接口可以实现多继承。
	// interface类型默认是一个指针（引用类型。）如果没有对interface初始化就是用 那么会输出nil
	// 空接口 interface{}没有任何方法，所以 所有类型的实现了空接口。也就是说：我们可以吧任何一个变量 赋给一个空接口类型。

	a.test01()
	a.test02()
	a.test03()

	var t T = s
	// t.test01() // 是错误的写法，  因为T是一个空接口。没有 定义任何方法。所以不能直接调用方法。
	fmt.Println(t) // {} 打印出来的 是一个空接口

	// 空接口变量 也可以这么写
	var t2 interface{} = s
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println("t2 = ", t2, ", t = ", t)

}
