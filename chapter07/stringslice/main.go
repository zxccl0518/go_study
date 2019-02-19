package main

import "fmt"

func main() {
	// strign底层是一个byte数组，因此string也可以进行切片处理。
	str := "hello world ,hello Go"
	var slice = str[6:11]
	fmt.Println("slice = ", slice) //slice =  world
	// string 不可以单独修改其中的某个字母。 如：str[0] = "k",编译不会通过，报错,原因是string是不可变的。
	// 如果需要修改字符串，先将字符串转换为[]byte切片，或者是修改为[]rune切片，然后修改，最后在重新转换为切片。
	// "hello world ,hello Go" => 改为
	var arr1 = []byte(str)
	arr1[0] = 'K'
	// arr1[0] = '北' //.\main.go:15:12: constant 21271 overflows byte
	str = string(arr1)
	fmt.Println("str = ", str)
	// 注意：我们转换成[]byte后，可以处理英文和数字，但是不能处理中文。
	// 原因是 []byte 字节来处理，而一个汉字是由3个字节组成的，因此会出现乱码。

	// 解决方式：将string转换为[]rune 即可，以为[]rune是按字符处理，兼容汉字。
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str = ", str)
}
