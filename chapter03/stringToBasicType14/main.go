package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str = "true"
	// var str = "helloOK"  helloOK不是 布尔类型，所以会将这个转换后的数据置为默认值，也就是false，其他类型也是一样的道理。【默认值】
	var b bool
	// b,_ = strconv.ParsBool()
	// 说明
	// 1.strconv.ParseBool(str) 函数会返回两个值(value bool, err error)
	// 2.因为我只想获取到value bool，不想获取err 所以我使用 _ 忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("\nb type:%T, b = %v \n", b, b)

	// ParseInt 返回的数据类型是 int64所以一定要用 int64接受，要想要int32， 类型装换一下就好
	var str2 = "1234590"
	var intValue int64
	intValue, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("intValue type:%T, intValue:%v \n", intValue, intValue)

	// 将int64的数据 转换为 int32
	var n1 int
	n1 = int(intValue)
	fmt.Printf("n1 type:%T, intValue:%v\n", n1, n1)

	// ParseFloat（str string, size bitsize） （float64, err error） 返回的数据类型是 float64
	var str3 = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type:%T, f1 : %v\n", f1, f1)

	// 将float64 装换位float32
	var f32 float32
	f32 = float32(f1)
	fmt.Printf("f_32 type:%T, f1 : %v\n", f32, f32)

	// 注意：将一个英文字母的  字符串 转换成int，因为没有对应的int类型， 所以是转换失败，结果是 0
	var str4 = "hello"
	var n3 int64
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type:%T, n3:%v \n", n3, n3)
	// 结果是：n3 type:int64, n3:0
}
