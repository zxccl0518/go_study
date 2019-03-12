package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"monsterName"`
	Age   int     `json:"monster_age"`
	Score float64 `json:"monster_score"`
	Sex   string
}

func (m Monster) GetSum(a int, b int) int {
	res := a + b

	return res
}

func (m Monster) Set(name string, age int, score float64, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func (m Monster) Print() {
	fmt.Printf("monster name = %v, age = %v, score = %v, sex = %v \n", m.Name, m.Age, m.Score, m.Sex)
}

func TestStruct(m interface{}) {
	rType := reflect.TypeOf(m)

	rVaL := reflect.ValueOf(m)

	// 想得到 m 有几个字段
	filedNum := rVaL.NumField()
	for i := 0; i < filedNum; i++ {
		if rType.Field(i).Tag.Get("json") != "" {
			fmt.Printf("filed index = %v, field[%v] = %v,  类型:%T \n", i, rType.Field(i).Tag.Get("json"), rVaL.Field(i), rVaL.Field(i))
		} else {
			fmt.Printf("filed index = %v, field[%v] = %v,  类型:%T \n", i, rType.Field(i).Name, rVaL.Field(i), rVaL.Field(i))
		}
	}

	// 想得到 m 有几个方法。
	methodNums := rVaL.NumMethod()
	fmt.Println("该 结构体有方法数目为 ：", methodNums)
	// 方法的排序 默认是按照函数名排序的。(ASCII)
	rVaL.Method(1).Call(nil) // 实际调用的是 print的方法。 方法按照ascii的排序 分别是 get  print set。 所以这里调用的是 print

	// 调用结构体第一个方法 Method(0)
	var params []reflect.Value // 声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(30))
	res := rVaL.Method(0).Call(params)
	fmt.Printf("res = %v,  res 类型 = %T \n", res[0].Int(), res)          //res = 40,  res 类型 = []reflect.Value  返回的切片是 value 类型的切片
	fmt.Printf("res = %v,  res 类型 = %T \n", res[0].Int(), res[0].Int()) //res = 40,  res 类型 = int64  将返回对应的 int int64 int32 int16 int8 类型的数据
}

func main() {
	monster := Monster{
		Name:  "牛魔王",
		Age:   100,
		Score: 99.9,
		Sex:   "male",
	}

	TestStruct(monster)
}
