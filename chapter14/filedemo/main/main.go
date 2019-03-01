package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	} else {
		fmt.Printf("file = %v\n", *file)

		file.Close()
	}
}
