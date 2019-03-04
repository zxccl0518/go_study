package testcase01

import (
	"testing"
)

// 编写要给测试用例，去测试AddUpper
func TestGetSub(t *testing.T) {
	// 调用
	res := getSub(10, 1)
	if res != 9 {
		t.Fatalf("getSub(10,1) 执行错误，期望值=%v  实际值=%v\n", 9, res)
	}

	t.Logf("getSub(10,1) 执行正确。。。")
}
