package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"monterName"`
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	// 演示
	monster := Monster{
		"牛魔王",
		100,
		"123123",
		1000.1,
		"牛魔拳",
	}

	// 将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败。err = ", err)
	}

	// 输出序列化后的json
	fmt.Println("序列化后的数据 data = ", string(data))
}

func main() {
	// 演示 将结构体，map，切片进行序列化。
	testStruct()

}
