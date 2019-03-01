package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令韩的参数有", len(os.Args))
	// 遍历 os.Args切片，就可以得到所有的命令行的参数。
	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v\n", i, v)
	}
}
