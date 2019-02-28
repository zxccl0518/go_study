package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "d:/test.txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件失败 err = ", err)
		return
	}
	// 及时关闭file句柄
	defer file.Close()

	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString("hello Gardon\r\n")
	}

	// 因为writer是带缓存的，隐刺在调用WriterString() 方法的时候，其实内容是先写到缓存中的，所以要调用 Flush()方法
	// 将缓冲区的内容写到文件中去，否则文件中会没有数据。
	writer.Flush()
}
