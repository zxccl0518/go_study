package main

import "fmt"

type Integer int

type Circle struct {
	Radios float64
	Area   float64
}

func (c *Circle) getArea() {
	fmt.Printf("getArea c = %p \n", c)
	c.Area = 3.14 * c.Radios // c.Area 等价于 （*c）.Area  底层会自动帮忙加 (*)
	c.Radios = 100
	fmt.Println("Area = ", c.Area, " Radios = ", c.Radios)
}

func (i *Integer) print() {
	*i = *i + 1
	fmt.Println("i = ", *i)
}

// 如果 一个method 实现了String() 方法 那么print 会输出string（）的内容
func (c *Circle) String() (newstr string) {
	newstr = fmt.Sprintf("Circle Area = %v, Radios = %v \n", c.Area, c.Radios)
	return
}

func main() {
	var i Integer = 1
	i.print()
	fmt.Println("i = ", i)

	c := Circle{10.0, 0}
	fmt.Println("Area = ", c.Area, " Radios = ", c.Radios)
	// (&c).getArea() //go 为了方便书写和理解，底层实现了 自动帮忙将指针 添加取地址符号& 和 取值符号*
	fmt.Printf("main c 地址是：%p \n", (&c))
	// 结果是 main中打印的地址 和 getArea()打印的地址是一致。
	(&c).getArea()
	fmt.Println("Area = ", c.Area, " Radios = ", c.Radios)

	fmt.Println("===========================================")
	fmt.Println("c = ", &c) //c =  Circle Area = 31.400000000000002, Radios = 100
	//func (c *Circle) String() (newstr string) 这个结构体实现了string的方法，但是要注意的是 参数中 传入的是指针，才可以 实现改动。参数 要对应。
}
