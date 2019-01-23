package main

import (
	"fmt"
	"unsafe"
)

// 演示golang中 + 的使用方式。
func main() {
	// *** 当 值超出了类型的取值范围，会报错，而不会转换成相应的负数值。

	// 有符号的 int8 8位 1个字节 相当于c byte
	var i int8 = -128
	// 有符号的 int16 16位 2个字节 相当于c short 其他的int32 int64以此类推。
	var int16Value int16 = 200

	fmt.Println("int8 i = ", i)
	fmt.Println("int16 i = ", int16Value)

	// 无符号的 uint8 1个字节， 0~255 其他的uint16，uint32，uint64 以此类推
	var unint8Value uint8 = 211
	fmt.Println("uint8 unint8Value = ", unint8Value)

	// --------------------------------------------------------------------------------

	// int 有符号，如果系统是32位的 那么占4个字节，
	// int 有符号，如果系统是32位的 那么占4个字节，

	// uint 无符号，如果系统是32位的 那么占4个字节，
	// uint 无符号，如果系统是64位的 那么占8个字节，

	// rune 有符号，与int32一样，  备注：等价于int32，表示一个Unicode码，
	// byte 无符号，与uint8等价    备注：当要存储字符时选择byte

	// int，uint，rune，byte 的使用。
	var a int = 8900
	fmt.Println("a = ", a)

	var b uint = 0
	var c byte = 255
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	// 整形的使用细节
	var n1 = 100
	// 这里我们给介绍一下如何查看某个变量的数据类型
	// fmt.Pringtf() 可以用于做格式化输出。
	fmt.Printf("n1 的类型是 %T \r\n", n1)

	//如何在程序查看某个变量的暂用字节大小和数据类型。(实用较多)
	var n2 int64 = 10
	// unsage.Sizeof(n2) 是unsage包的一个函数，可以返回n2变量占用的字节数。
	fmt.Printf("n2 的数据类型: %T n2占用的字节数是:%d ", n1, unsafe.Sizeof(n2))

}
