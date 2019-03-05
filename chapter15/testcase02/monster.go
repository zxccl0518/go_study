package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(" 序列化失败 err = ", err)
		return false
	}

	filePath := "d:/monster.txt"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("文件写入失败 err = ", err)
		return false
	}
	return true
}

func (m *Monster) Restore() bool {
	// 1.先从文件中读取 序列化的字符串。
	filePath := "d:/monster.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败 err = ", err)
		return false
	}

	//2.使用读取到的data[]byte 反序列化
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("反列化失败 err = ", err)
		return false
	}

	return true
}
