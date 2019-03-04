package testcase01

import (
	"fmt"
	"testing"
)

// 编写要给测试用例，去测试AddUpper
func TestAddUpper(t *testing.T) {
	// 调用
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v  实际值=%v\n", 55, res)
	}

	t.Logf("AddUpper(10) 执行正确。。。")
}

func TestHello(t *testing.T) {
	fmt.Println("测试 Testhello ...")
}
