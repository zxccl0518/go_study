package main

import (
	"fmt"
	"testing"
)

func testPrint(t *testing.T) {
	res := Print1To20()

	// 这个函数是 跳过 test,不论后面的结果 是对还是错的,都会跳过.
	// t.SkipNow()

	if res != 210 { // 210是随机写的一个数字.
		t.Errorf("wrong result ---------------res = %v\n", res)
	} else {
		fmt.Println(" 测试结果通过 恭喜 结果 = ", res)
	}
}

// 为了验证 subtest  ,t.Run来执行,subtest额快充做到控制test输出以及test的顺序.
func testSubTest(t *testing.T) {
	t.Run("a1", func(t *testing.T) { fmt.Println("a1") })
	t.Run("a2", func(t *testing.T) { fmt.Println("a2") })
	t.Run("a3", func(t *testing.T) { fmt.Println("a3") })
}

// TestMain(m *testing.M){}
// 使用TestMain 作为初始化test,并且使用m.Run()来调用其他tests 可以完成一些需要初始化操作的testing,比如数据库连接,文件打开,rest服务登录灯.
// 如果没有再TestMainvs调用m.Run()则除了TestMain以外的其他的Tests都不会被执行.
func testMain(m *testing.M) {
	fmt.Println("test main begins")

	// TestMain 如果不掉用m.Run就不会运行其他的test case了.
	m.Run()
}

// 综合测试
func TestAll(t *testing.T) {
	// 这里的testPrint改为小写的原因是 让其成为test的子程序,大写的话 会单独执行.  就会多执行一遍.
	t.Run("testPrint", testPrint)
	t.Run("testSubTest", testSubTest)
}

// 以下为TestAll的测试结果.
// 测试结果通过 恭喜 结果 =  210
// a1
// a2
// a3

func BenchmarkAll(b *testing.B) {
	fmt.Println(" N = ", b.N)
	for n := 0; n < b.N; n++ {
		Print1To20()
	}
}
