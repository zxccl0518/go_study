package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型转换为string类型

func main() {
	var num1 = 99
	var num2 = 23.456
	var b = true
	var myChar byte = 'h'
	var str string // 空的str

	// 使用第一种方式来转换， fmt.Sprintf方法。
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("\nnum1 -> str type: %T , str = %q \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("num2 -> str type: %T , str = %q \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("b -> str type: %T , str = %q \n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("myChar -> str type: %T , str = %q \n", str, str)

	// 方式2 strconv包函数
	// func FormatBool(b bool)
	// func FormatInt()
	//
	//
	var num3 = 99
	var num4 = 23.456
	var b2 = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("\n\n\nnum3 -> str type: %T , str = %q \n", str, str)

	// strconv.FormatFloat(num4, 'f', 10, 64)
	// 说明：'f':格式， 10：表示小数位保留10位， 64:表示这个描述是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("num4 -> str type: %T , str = %q \n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("b2 -> str type: %T , str = %q \n", str, str)

	// 将 int 直接转换为 string, 如果是int64，需要强制类型转换。
	var num5 = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("num5 itoa() -> str type: %T , str = %q \n", str, str)

	var num6 int64 = 78911
	str = strconv.Itoa(int(num6))
	fmt.Printf("num6 itoa() int64Toint -> str type: %T , str = %q \n", str, str)

}
