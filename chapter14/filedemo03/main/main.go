package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 使用ioutil.ReadFile 一次性将文件读取到到位
	file := "d:/test.txt"
	content, err := ioutil.ReadFile(file) // 采用这个方式的主要原因是 因为文件小，如果文件很大的话，这种的方式就会很耗时，不能采用这种方式。
	if err != nil {
		fmt.Println("read file err = ", err)
	}
	// 把读取到的内容 显示到终端。
	// fmt.Printf("%v\n", content)         //[]byte // 打印出来的都是数字，因为content是byte[]
	fmt.Printf("%v\n", string(content)) //强制类型 转换只会 就会正常的显示汉字了。
	// 我们没有显示的Open文件，因此也不需要显示的Close()
	// 因为，文件的Open 和Close被封装到ReadFile 函数内部了.
}
