package monster

import "testing"

func TestStore(t *testing.T) {
	m1 := &Monster{
		"猪八戒",
		100,
		"吃的贼多",
	}
	res := m1.Store()

	if res == false {
		t.Fatalf("测试 Store 函数失败")
	}

	t.Logf("测试 Store() 成功")
}

func TestRestore(t *testing.T) {
	// 测试数据是很多， 测试很多次才确定函数 模块的稳定性
	var m Monster

	res := m.Restore()
	if res == false {
		t.Fatalf("测试 Restore() 方法失败.")
	}

	t.Logf("测试反序列化 成功 Name = %v Age = %v Skill = %v \n", m.Name, m.Age, m.Skill)
}
