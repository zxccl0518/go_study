package main

import "fmt"

func main() {
	// 输出10局 你好尚硅谷
	// golang 中，有循环控制语句，来处理循环的执行某段代码的方法 -> for 循环。

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(" 你好 尚硅谷。")
	// }

	// 也可以这么写，把变量定义在for之前。
	// i := 1
	// for ; i <= 5; i++ {
	// 	fmt.Println("第二种方式，打印for 循环。")
	// }

	// for的第2种写法。 迭代。
	// j := 1
	// for j <= 3 {
	// 	fmt.Println("for 的第三种写法。。 打印 你好 尚硅谷 ~")
	// 	j++
	// }

	// for的第三种写法。 通常需要break 跳出死循环。 这种写法等价于 for ; ; {}
	// k := 1
	// for {
	// 	fmt.Println("ok ~~~")
	// 	k++

	// 	if k > 2 {
	// 		break
	// 	}
	// }

	// 字符串中有汉字，这种传统的便利方法 是按照字节去取值的。 因为golang 是utf-8的编码格式，中文是3个字节，所以会出现乱码
	// 那如何解决呢，将str 转换为 rune切片。
	// str := "hello world, hello 北京"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c", str[i])
	// }

	// 切片的方式 解决乱码的问题。
	var str = "hello，world 北京！"
	str2 := []rune(str) // []rune 切片。
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}

	// for 第四种方式。 for-range 可以方便 遍历字符串和数组
	// 这种迭代器的方式，不会出现乱码。原因是 迭代器是按照字符的方式来遍历的。遇到了汉字 index索引会跳动3个。表示3个字节。
	// 	str := "hello world, hello! 北京。"
	// 	for index, val := range str {
	// 		fmt.Printf("index = %d val = %c \n", index, val)
	// 	}
}
