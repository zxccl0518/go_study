package globalutils

import "fmt"

var GlobalValue = "this is global value"

func Test() {
	fmt.Println("globalutils test")
}

func init() {
	fmt.Println("globalutils 包的 init 函数。")
}
