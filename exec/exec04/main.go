package main

import "fmt"

func main() {
	// 1) 使用switch 吧小写的类型char 型，转换为大写的(键盘输入。)
	// 2） 制转换a,b,c,d,r, 其他的输出other

	// var char byte
	// fmt.Println("请输入一个 字符")
	// fmt.Scanf("%c", &char)
	// switch char {
	// case 'a':
	// 	fmt.Printf("char 的大写字母是:%c \n", 'a'-31)
	// case 'b':
	// 	fmt.Printf("char 的大写字母是: %c\n", 'a'-31)
	// default:
	// 	fmt.Println("a 是其他值")
	// }

	// 2）对学生成绩大于60分的，数车"合格"。 低于60分的，输出"不合格"
	// 注：输入的成绩不能大于100
	// var score float64
	// fmt.Println("请输入成绩：")
	// fmt.Scanf("%f", &score)
	// switch {
	// case score < 60:
	// 	fmt.Println("成绩低于60 不及格")
	// case score <= 100 && score > 60:
	// 	fmt.Println("成绩大于 60 且 小于100 成绩为 合格")
	// default:
	// 	fmt.Println(" 您输入的成绩 已经大于了 100的上线。 输入有误。")
	// }
	// 这段也可以换一种写法
	// var score float64
	// fmt.Println("请输入成绩")
	// fmt.Scan(&score)
	// switch int(score / 60) {
	// case 1:
	// 	fmt.Println("及格了")
	// case 0:
	// 	fmt.Println("不及格")
	// default:
	// 	fmt.Println("输入成绩无效")
	// }

	// 3）根据用户的指定月份，
	// 打印盖玉芬所属的季节，345 为春季，678 为夏季，9 10 11 为秋季，12 1 2为冬季
	var month int
	fmt.Println("请输入 月份：")
	fmt.Scan(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println("该 月份为春季。")
	case 6, 7, 8:
		fmt.Println("该月份为 夏季 ")
	case 9, 10, 11:
		fmt.Println("该月份为 秋季 ")
	default:
		fmt.Println("该月份为 冬季。")
	}

	// 如果判断的具体数值不多，而且符合整数，浮点数，字符，字符串，这几种类型，建议使用switch语句，简洁高效。
	// 其他情况，对 区间判断 和结果为 bool类型的判断，使用if， if的适用范围 更广泛。
}
