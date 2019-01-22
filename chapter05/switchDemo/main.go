package main

import "fmt"

func main() {
	// switch 后也可以不带表达式，类似if --else 分支来使用。【案例演示。】
	var age = 10
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配到。")
	}

	// case 中也可以对范围进行判断。
	var score = 99
	switch {
	case score > 90:
		fmt.Println("成绩优秀。")
	case score > 70 && score <= 90:
		fmt.Println("成绩优良。")
	case score >= 60 && score <= 70:
		fmt.Println("成绩及格。")
	default:
		fmt.Println("不及格。")
	}

	// switch后也可以直接声明/定义 一个变量，分号结束，但是不推荐这种使用方法。[案例演示]
	switch grade := 67; {
	case grade > 90:
		fmt.Println("成绩优秀。")
	case grade > 70 && grade <= 90:
		fmt.Println("成绩优良。")
	case grade >= 60 && grade <= 70:
		fmt.Println("成绩及格。")
	default:
		fmt.Println("不及格。")
	}

	// switch 穿透的使用方法。 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		// 匹配到了case 10的情况，执行完成case10  会继续向下执行case20的代码。默认只有穿透了1层。
		fallthrough
	case 20:
		fmt.Println("ok2")
	default:
		fmt.Println("没有匹配成功")
	}

	// type switch : switch语句还可以被用于type-switch 来判断某个interface变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Println("x的类型 ： 空值")
	case int:
		fmt.Println("x的类型 是 int型")
	case float64:
		fmt.Println("x的类型是 float64型， 值：%v ", i)

	}
}
