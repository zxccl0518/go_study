package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Student struct {
	Name string
	Age  int
}

type Stu Student

type Integer int

// type Monstor struct {
// 	Name  string
// 	Age   int
// 	Skill string
// }
type Monstor struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	fmt.Println(a, b)

	// a = b    // .\main.go:18:4: cannot use b (type B) as type A in assignment 报错。
	a = A(b) // 通过强制类型转换，成功了，原因是因为2个结构体拥有 完全相同的字段。包括名字、个数、类型

	fmt.Println(a, b)

	// 结构体记性type重新定义(相当于区别名)，golang认为是新的数据类型，但是可以互相相互强制类型转换。
	var stu1 Student
	var stu2 Stu
	// stu2 = stu1 // 错误的操作，如果想实现，可以通过强制类型转换 来实现。
	stu2 = Stu(stu1) // 错误的操作，如果想实现，可以通过强制类型转换 来实现。
	fmt.Println(stu1, stu2)

	// 类似的，其他的类型的变量， 类型一旦b被type重新定义了，就不能直接赋值了。
	var i1 int
	var i2 Integer

	// 是个错误的操作。
	// i1 = i2 //.\main.go:43:5: cannot use i1 (type int) as type Integer in assignment
	// 但是可以通过 强制类型转换，来实现。
	i1 = int(i2)
	fmt.Println(i1, i2)

	// 结构体重Tag的使用 演示实例
	// 1.创建一个Monster比那辆
	monster := Monstor{"牛魔王", 100, "杀人"}
	// 2. 将monster变量序列化为一个json格式字串
	// Marshal 函数返回值有2个，第一个是[]byte, 第二个是错误信息。
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误 ", err)
	}
	// fmt.Println("jsonStr = ", jsonStr) //sonStr =  [123 34 78 97 109 101 34 58 34 231 137 155 233 173 148 231 142 139 34 44 34 65 103 101 34 58 49 48 48 44 34 83 107 105 108 108 34 58 34 230 157 128 228 186 186 34 125]
	fmt.Println("jsonStr = ", string(jsonStr)) //jsonStr =  {"Name":"牛魔王","Age":100,"Skill":"杀人"} ，所有的字段 首字母都是大写的。若将结构体中首字母 小写了，那么json.Marshal(monster)这个 执行结果就为空，因为字母小写了，那么其他包无法访问这个变量。所以 这个时候tag 就有用了。
	// 当结构体中 使用了tag标签之后，解析出来的字段都是小写的了，而且还不会影响其他包 调用本包的结构体变量。
	// jsonStr =  {"name":"牛魔王","age":100,"skill":"杀人"}

}
