package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json : "nameMonster"`
	Age      int
	Birthday string //....
	Sal      float64
	Skill    string
}

// 演示对反序列化 将json反序列化成 结构体。
func unmarshalStruct() {
	// str 在项目开发中，是通过网络传输获取到的，或者是通过读取文件获取到的。
	str := "{\"Name\":\"牛魔王\",\"Age\":100,\"Birthday\":\"123123\",\"Sal\":1000.1 , \"Skill\":\"牛魔拳\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化失敗。 err = ", err)
		return
	}
	fmt.Printf("Name:%v\t age:%v\t birthday:%v\t sal:%v\t skill:%v\n", monster.Name, monster.Age, monster.Birthday, monster.Sal, monster.Skill)
}

// 反序列化 map 演示。
func unmashalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"

	// 定义一个map 反序列化不需要给map make
	var a map[string]interface{}

	// 反序列化
	// 反序列化 map不需要make的原因是，make操作被封装到Unmashal函数中了。
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(" 反序列化 失败。err = ", err)
		return
	}
	fmt.Println("反序列化后的 map a = ", a)
}

// 演示 将json字符串反序列化成切片。
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"

	var s []map[string]interface{}

	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(" 反序列化 失败。err = ", err)
		return
	}
	fmt.Println("反序列化后的 slice = ", s)
}

func main() {
	unmarshalStruct()

	unmashalMap()

	unmarshalSlice()
}
