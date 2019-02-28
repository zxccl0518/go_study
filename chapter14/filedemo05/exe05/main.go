package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 将d:/ test.txt 文件内容导入到 d：aaa.txt中

	// 1. 首先将d:/test.txt的内容读取到内存，
	// 2. 将读取到的内容 写入到 d:/aaa.txt中

	filePath1 := "d:/test.txt"
	filePath2 := "d:/aaa.txt"

	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("readFile err = ", err)
		return
	}

	// str := fileSrc.
	err = ioutil.WriteFile(filePath2, data, 0666)
	if err != nil {
		fmt.Println("写文件 出错 err = ", err)
	}
}
