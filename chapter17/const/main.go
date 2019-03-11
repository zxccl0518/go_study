package main

import "fmt"

func main() {
	var num int
	num = 9

	// const tax int	// const declaration cannot have type without expression 一定要赋值 初始化
	const tax int = 1
	fmt.Println(num, tax)
	// 常量不能修改
	// 常量只能修饰 int float bool string

	// num1 := num / 3
	// const b = num / 3 // 报错 因为num 是可变的。
	// const b = num1   // 也是错的
	// const c = getVal() // 也是错的

	// const (
	// 	a    = iota       //表示 啊默认是0 后面的值是递增的			  0
	// 	b    = iota       // 1
	// 	c, d = iota, iota // 如果这么写的话，就是表示这2个值是一样的，不会递增 // 2,2
	// 	e    = iota       // 3
	// )

	const (
		a = 9
		b = iota
		c
		d
		e
	)

	fmt.Println(a, b, c, d, e)
}
