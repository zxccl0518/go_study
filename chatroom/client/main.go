package main

import "fmt"

var userID int
var userPwd string
var userName string

func main() {
	// 接收 用户的选择。
	var key int
	for {
		fmt.Println("-----------------------欢迎登录多人聊天室-----------------------")
		fmt.Println("-----------------------1.登录聊天室----------------------------")
		fmt.Println("-----------------------2.注册用户----------------------------")
		fmt.Println("-----------------------3.退出系统----------------------------")
		fmt.Println("-----------------------请选择（1~3）:----------------------------")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入 id：")
			fmt.Scanln(&userID)
			fmt.Println("请输入 密码：")
			fmt.Scanln(&userPwd)
			err := login(userID, userPwd)
			if err != nil {
				fmt.Println("登录失败")
			} else {
				fmt.Println("登录成功")
			}
		case 2:
		case 3:
		default:
		}
	}
}
