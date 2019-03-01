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

// 判断文件存在或者是文件夹是否存在。
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func main() {
	// 将 d：test。txt 文件拷贝到 d:/QMDownload/中
	srcFile := "e:/abc.txt"
	destFile := "e:/456.txt"

	len, err := copy(destFile, srcFile)

	if err != nil {
		fmt.Println("拷贝失败 err = ", err)
	} else {
		fmt.Println("拷贝成功 len = ", len)
	}

	isExist, err := PathExists(srcFile)
	if isExist == true {
		fmt.Println(srcFile, "文件 是存在的。")
	} else {
		if err == nil {
			fmt.Println(srcFile, "文件 是 不存在的。")
		} else {
			fmt.Println(srcFile, "文件 是 不存在的。 err = ", err)
		}
	}
}
