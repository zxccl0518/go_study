package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

//
func testStruct() {
	// 演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	// 将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
	}

	// 输出序列化后的结果
	fmt.Printf("monster 序列化后后 = %v \n", string(data))
}

// 将map序列化
func testMap() {
	// 定义一个map
	var m map[string]interface{}
	m = make(map[string]interface{})

	m["name"] = "红孩儿"
	m["age"] = 30
	m["address"] = "洪崖洞"

	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
	}

	// 输出序列化后的结果
	fmt.Printf("m map 序列化后后 = %v \n", string(data))
}

// 演示对切片进行序列化，我们这个且前 []map[string] interface{}
func testSlice() {
	var s []map[string]interface{}
	s = make([]map[string]interface{}, 0)
	var m1 = make(map[string]interface{})
	m1["1"] = 123
	m1["2"] = "321"
	s = append(s, m1)

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("序列化失败 err = %v\n", err)
		return
	}
	fmt.Println("序列化后的data = ", string(data))
}

func main() {
	// 静结构体 map 切片进行序列化。
	// testStruct()
	testMap()

	testSlice()
}
