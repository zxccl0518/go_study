package main

import "fmt"

func main() {
	// 加入还有97天放假，问：xx个星期零xx天？
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d个星期零%d天", week, day)

	// 定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：
	// 5/9*(华氏温度-100)，请求出华氏温度对应的摄氏温度。
	var huashi float32 = 123.4
	// var sheshi float32 = 5 / 9 * (huashi - 100) // 这里要注意一下， 如果写得是5 则结果就是为0.000000， 因为5/9 = 0。 但是写成5.0， 就是 5.0/9 结果是有小数的。
	var sheshi float32 = 5.0 / 9 * (huashi - 100) // 这里要注意一下， 如果写得是5 则结果就是为0.000000， 因为5/9 = 0。 但是写成5.0， 就是 5.0/9 结果是有小数的。
	fmt.Printf("华氏温度为123.4, 对应的摄氏温度为%f", sheshi)  //13.000002
}
