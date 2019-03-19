package main

import "fmt"

func main() {
	var key int
	var userID string
	var userPwd string

	for {
		fmt.Println("-------------------欢迎登录 聊天房间系统 ---------------------")
		fmt.Println("\t\t\t 1.登录聊天房间\t\t\t")
		fmt.Println("\t\t\t 2.注册用户\t\t\t")
		fmt.Println("\t\t\t 3.退出聊天系统\t\t\t")
		fmt.Println("请输入1~3:")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入 用户id:")
			fmt.Scanf("%s\n", &userID)
			fmt.Println("请输入 用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			err := login(userID, userPwd)
			if err != nil {
				fmt.Println("登录失败 err = ", err)
			} else {
				fmt.Println("登录成功 err = ", err)
			}
		case 2:
		case 3:
		default:
		}
	}
}
