package main

import (
	"fmt"
	"os"

	"github.com/zxccl0518/go_study/chatroom/client/process"
)

func main() {
	var key int
	var userID int
	var userPwd string
	var usrName string
	// var loop = true
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
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入 用户密码：")
			fmt.Scanf("%s\n", &userPwd)

			// 创建结构体，完成登录的功能。
			up := &process.UserProcess{}
			up.Login(userID, userPwd)
			// loop = false
		case 2:
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userID)
			fmt.Printf("请输入 用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("%s\n")
			fmt.Println("请输入用户的名字：")
			fmt.Scanf("%s\n", &usrName)
			up := &process.Rigister
			up.Rigister(userID, userPwd, usrName)
			// loop = false
		case 3:
			os.Exit(0)
			break
		default:
		}
	}
}
