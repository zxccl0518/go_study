package main

import (
	"errors"
	"fmt"
)

// func test() {
// 	num1 := 10
// 	num2 := 0
// 	res := num1 / num2 // 这里是有问题的，因为除数不能为0
// 	fmt.Println("res = ", res)
// }

// 上面的错误代码的 解决方案。
func test() {
	// 使用defer + recover 来捕获和处理异常。
	defer func() {
		err := recover() // recover()是内置函数，可以捕获到异常。
		if err != nil {  // 说明捕获到错误。
			// if err:=recover; err != nil{			// 这个写法 等同于上面2行的写法。

			fmt.Printf("err信息 = %v， err 类型=%T\n", err, err)
			// 这里就可以将错误信息 发送给管理员。
			fmt.Println("发送邮件给管理员。")
		}
	}()
	// 通过这种方式，可以在捕获异常的情况下，让程序继续运行。

	num1 := 10
	num2 := 0
	res := num1 / num2 // 这里是有问题的，因为除数不能为0
	fmt.Println("res = ", res)
}

// 函数去读取配置文件init.conf的信息。
// 如果文件名传入不正确，我们就返回一个自定义的错误。
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取。。。
		return nil
	} else {
		// 返回一个自定义错误。
		return errors.New("读取文件错误。")
	}
}
func test02() {
	// err := readConf("config.ini")
	err := readConf("con.ini")
	if err != nil {
		// 如果 读取文件发生错误，就输出这个错误，并终止程序。
		panic(err)
	}

	fmt.Println("test02() 继续执行。。。")
}

func main() {
	// 测试
	// test() //运行时报错：panic: runtime error: integer divide by zero
	// for {
	// 	fmt.Println("main()中下面的代码 。。。")
	// 	time.Sleep(time.Second)
	// }

	test02()
}
