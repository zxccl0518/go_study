package process

import (
	"fmt"
	"os"
)

// 显示登录成功后的界面。
func showMenu() {
	fmt.Println("--------- 恭喜 xxx 登录成功 --------")
	fmt.Println("\t\t\t 1.显示在线用户列表\t\t\t")
	fmt.Println("\t\t\t 2.发送消息。\t\t\t")
	fmt.Println("\t\t\t 3.信息列表。\t\t\t")
	fmt.Println("\t\t\t 4.提出系统。\t\t\t")
	fmt.Println("\t\t\t 请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示 在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你 退出了体统")
		os.Exit(0)
	default:
	}

}
