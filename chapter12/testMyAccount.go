package main

import "fmt"

func main() {
	// 声明一个变量，保存接受用户输入选项。
	key := ""
	// 声明一个变量，控制是否退出for
	loop := true
	// 声明一个变量，保存账户余额
	balance := 10000.0
	// 每次收支金额
	money := 0.0
	// 每次收支说明
	note := ""
	// 收支的详情，使用字符串来记录
	details := "收支\t账户金额\t收支金额\t 说	明"
	// 显示这个主菜单。
	for {
		fmt.Println("--------------------------家庭收支记账软件----------------------------")
		fmt.Println("                          1.收支明细")
		fmt.Println("                          1.登记收入")
		fmt.Println("                          1.登记支出")
		fmt.Println("                          1.退出软件")
		fmt.Println("请选择(1-4):")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("--------------------------当前收支明细-------------------------")
			fmt.Println(details)
		case "2":
			fmt.Println("--------------------------登记收入----------------------------")
			fmt.Print("本次收入 金额：")
			fmt.Scanf("%f \n", &money)
			fmt.Print("本次收入 说明：")
			fmt.Scanf("%s \n", &note)
			balance += money

			details += fmt.Sprintf("\n收入\t%v\t%v\t%v\n", balance, money, note)
			fmt.Println(details)
		case "3":
			fmt.Println("--------------------------登记支出----------------------------")
		case "4":
			loop = false
		default:

		}

		if !loop {
			break
		}
	}
	fmt.Println("你退出家庭记账软件的使用。")
}
