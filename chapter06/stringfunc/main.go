package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1 统计字符串的长度，安子杰len（str）
	str := "hello北" // golang 的编码同一位utf-8（ascii的字符(字母和数字)占一个字节，汉字占用3个字节。）
	fmt.Println("str len = ", len(str))

	//2  字符串遍历，[]rune(str) 采用切片的方式，保证不会乱码。
	str2 := "hello北京"
	strrune := []rune(str2)
	for i := 0; i < len(strrune); i++ {
		// fmt.Println("字符 = ", str2[i])	// 打印出胡来的都是数字
		// fmt.Printf("字符 = %c \n", str2[i]) // 打印出胡来的都是字母，中文部分乱码。
		fmt.Printf("字符 = %c len = %d \n", strrune[i], len(strrune)) // 打印出胡来的都是字母，中文可以正常显示。
	}

	//3 字符串转整数：n,err := strconv.Atoi("21")
	// n, err := strconv.Atoi("123")	// 结果是n=123
	n, err := strconv.Atoi("hellp") // 结果是strconv.Atoi: parsing "hellp": invalid syntax
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("n = ", n)
	}

	//4 整数转字符串。 str := strconv.Itoa(123)
	stritoa := strconv.Itoa(123)
	fmt.Printf("stritoa = %v, stritoa = %T \n", stritoa, stritoa)

	//5 字符串转[]byte var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes = %v\n", bytes) //[104 101 108 108 111 32 103 111] ascii 的形式打印出来。

	//6 []byte 转字符串： str = string([]byte{97,98,99})
	btoaStr := string([]byte{97, 98, 99})
	fmt.Printf("btoastr = %v\n", btoaStr) //btoastr = abc

	//7  10进制转换 2,8,16进制： str = strconv.FormatInt(123,2)
	str = strconv.FormatInt(123, 2)               //十进制转换为2进制。
	fmt.Printf("str = %v, str = %T \n", str, str) //str = 1111011, str = string
	str = strconv.FormatInt(123, 8)               //十进制转换为8进制。
	fmt.Printf("str = %v, str = %T \n", str, str) //str = 1111011, str = string
	str = strconv.FormatInt(123, 16)              //十进制转换为16进制。
	fmt.Printf("str = %v, str = %T \n", str, str) //str = 1111011, str = string

	// 8 查找子串 是否在指定的字符串中，strings.Contains(s, substring string) bool
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b = %v, b = %T \n", b, b) //b = true, b = bool
	b = strings.Contains("seafood", "m")
	fmt.Printf("b = %v, b = %T \n", b, b) //b = false, b = bool

	//9 统计一个字符串中有几个指定的子串  strings.Count(s, seedstr string) int
	num := strings.Count("chinese", "e")
	fmt.Println("num = ", num) // 查找到了2个

	//10  不区分大小写的字符串比较。(== 是区分大小写的) strings.EqualFold("abc", ABC) bool
	b = strings.EqualFold("abc", "ABC")
	fmt.Println("b = ", b)              // true
	fmt.Println("b = ", "abc" == "ABC") // false 因为 == 区分大小写。

	// 11)返回子串在字符串[第一次]出现的index值，如果没有返回-1
	// strings.index("nLT_abc", "abc)
	index := strings.Index("nLT_abc", "abc")
	fmt.Println("index= ", index) //4
	index = strings.Index("nLT_abc", "hello")
	fmt.Println("index= ", index) // -1
	index = strings.Index("abc_abc", "abc")
	fmt.Println("index= ", index) // 2

	//12） 返回子串在字符串最后一次出现的index，如果没有返回-1， strings.LastIndex("go golang", "go")
	index = strings.LastIndex("go go go golang", "go")
	fmt.Println("index = ", index) // index = 9  因为index 是从0开始的。

	// 13) 将制定的子串替换成另外一个子串， strings.Replace("go go hello", "go", apple, n) n可以指定你希望替换几个，如果n = -1表示替换全部。
	str = strings.Replace("go go hello", "go", "apple", 1)
	fmt.Println("strings.Replace n = 1, str = ", str)
	str = strings.Replace("go go hello", "go", "apple", -1)
	fmt.Println("strings.Replace n = -1, str = ", str)

	//14) 按照指定的某个字符，为分割标识，将一个字符串拆分成 字符串数组。 strings.Split("hello,world,ok", ",")
	strsplit := "hello, world, apple"
	strArr := strings.Split(strsplit, ",")
	fmt.Println("strings.Split 原 strsplit = ", strsplit) // 原来的字符串是不会 被改变的。
	fmt.Println("strings.Split 新 strArr = ", strArr)

	//15 将字符串的字母进行大小写的转换。：strings.ToLower("Go") // go.  string.ToUpper("Go") -> GO
	str = "AaAa"
	strToLower := strings.ToLower(str)
	fmt.Println("strings.ToLower strToLower = ", strToLower) //aaaa
	strToUpper := strings.ToUpper(str)
	fmt.Println("strings.ToLower ToUpper = ", strToUpper) //AAAA

	//16) 将字符串左右两边的空格去掉。 strings.TrimSpace("aa bb cc dd")
	str = " aa bb cc dd "
	strTrim := strings.TrimSpace(str)
	fmt.Printf("TrimSpace --- strTrim =%q", strTrim) // "aa bb cc dd" 可以看见 左右边的空格被去掉了

	//17) 将字符串左右两边的指定的字符去掉。
	// strings.Trim("! hel!lo! ", " !")		// ["hello"]// 将左右两边的 ! 和 " " 去掉。
	strTrim = strings.Trim("! hel!lo! ", " !")
	fmt.Printf("strings.Trim --- strTrim =%q\n", strTrim) //"hel!lo" 中间的!没有去掉。只去掉左右两边的。

	//18)、19) 还有 strings.TrimLeft() 和 strings.TrimRight()用法同strings.Trim() 一个是仅仅去掉字符串左边的 字符，一个是去掉右边的字符。

	// 20)判断字符串是否以指定的字符串开头:strings.HasPrefix("ftp:192.168.10.1","ftp")
	b = strings.HasPrefix("ftp:192.168.10.1", "ftp")
	fmt.Println("strings.HasPrefix  b = ", b)
	b = strings.HasPrefix("ftp:192.168.10.1", "http")
	fmt.Println("strings.HasPrefix  b = ", b)

	// 21) 判断字符串是否已指定的后缀字符串 结尾。 strings.HasSuffix("NLT_abc.png","png")
	// 用法同上面的 strings.HasPrefix() 在这里就不举例子了。
}
