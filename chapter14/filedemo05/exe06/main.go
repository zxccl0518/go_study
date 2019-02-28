package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 接受2个文件路径，一个是srcFileName 另一个是destFileName

func copy(destFileName string, srcFileName string) (written int64, err error) {
	file, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("打开文件出错 err = ", err)
	}
	defer file.Close()

	// 通过file 获取到Reader
	reader := bufio.NewReader(file)

	destFile, err := os.OpenFile(destFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer destFile.Close()

	writer := bufio.NewWriter(destFile)
	return io.Copy(writer, reader)
}

func main() {
	// 将 d：test。txt 文件拷贝到 d:/QMDownload/中
	srcFile := "d:/123.jpg"
	destFile := "d:/QMDownload/456.jpg"

	len, err := copy(destFile, srcFile)

	if err != nil {
		fmt.Println("拷贝失败")
	} else {
		fmt.Println("拷贝成功 len = ", len)
	}
}
