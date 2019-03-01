package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件//
	// 概念说明 file的叫法
	// 1. file 叫file对象
	// 2. file 叫file指针
	// 3. file 叫file文件句柄

	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}

	// 当函数退出的时候，要及时关闭file
	defer file.Close()

	// 创建一个*Reader, 是带缓冲垫额。
	/*
		const(
			defaultBufSize = 4096 // 默认的缓冲区为 4096
		)
	*/
	reader := bufio.NewReader(file) // 一点一点的读取
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 代表 文件读到了末尾了。
			break
		}

		fmt.Print("str = ", str)
	}
	fmt.Println("文件 读取结束。")
}
