package main

import (
	"fmt"

	"github.com/zxccl0518/go_study/customerManager/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

func (this *customerView) list() {
	this.customerService.List()
}

// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件--------------------")
		fmt.Println("                 1.添加客户                         ")
		fmt.Println("                 2.修改客户                         ")
		fmt.Println("                 3.删除客户                         ")
		fmt.Println("                 4.客户列表                         ")
		fmt.Println("                 5.退出                             ")
		fmt.Print("请选择(1-5)")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
			this.list()
		case "5":
			fmt.Println("退出")
			this.loop = false
		default:
			fmt.Println("您输入有误。")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("你 退出了客户关系管理系统。")
}

func main() {

	//创建一个customerView实例，并运行主菜单。
	cv := customerView{
		"",
		true,
		service.NewCustomerService(),
	}

	cv.mainMenu()
}
