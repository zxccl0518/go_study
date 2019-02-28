package main

import (
	"fmt"

	"github.com/zxccl0518/go_study/customerManager/model"
	"github.com/zxccl0518/go_study/customerManager/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

// 展示 用户列表信息
func (this *customerView) list() {
	customers := this.customerService.List()

	fmt.Println("------------------------------客户信息----------------------------------")
	fmt.Println("Id\tName\tGender\tAge\tPhone\tEmail")
	for i := 0; i < len(customers); i++ {
		str := customers[i].GetInfo()
		fmt.Println(str)
	}
	fmt.Println("------------------------------客户信息展示完毕---------------------------")
}

//增加一个新用户
func (this *customerView) add() {
	fmt.Print("名字：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("电邮：")
	email := ""
	fmt.Scanln(&email)

	cus := model.NewCustomer1(name, gender, age, phone, email)
	this.customerService.AddCustomer(cus)
}

// 删除用户
func (this *customerView) delete() {
	fmt.Print("请输入 要删除用户的id：")
	var id = 0
	fmt.Scanln(&id)
	res := this.customerService.FindById(id)
	if res == -1 {
		fmt.Println("输入的用户id 有误， 删除失败")
	} else {
		fmt.Print("是否确定 删除用户 确认请按y，取消请按n\n")
		var str string
		fmt.Scanln(&str)
		if str == "y" {
			this.customerService.DeleteCustomer(res)
		} else {
			fmt.Println("取消删除 成功")
		}
	}
}

// 修改客户信息。
func (this *customerView) update() {
	fmt.Print("请输入 要修改客户的id:")
	var numId = 0
	fmt.Scanln(&numId)
	res := this.customerService.FindById(numId)
	if res == -1 {
		fmt.Println("您 输入的用户的id有误，修改用户信息失败。")
	} else {
		customers := this.customerService.List()
		customerInfo := customers[res].GetInfo()
		fmt.Println("Id\tName\tGender\tAge\tPhone\tEmail")
		fmt.Println(customerInfo)
		fmt.Println("请输入 新的姓名:")
		var name = ""
		fmt.Scanln(&name)
		fmt.Println("请输入 新的性别:")
		var gender = ""
		fmt.Scanln(&gender)
		fmt.Println("请输入 新的年龄:")
		var age = 0
		fmt.Scanln(&age)
		fmt.Println("请输入 新的电话:")
		var phone = ""
		fmt.Scanln(&phone)
		fmt.Println("请输入 新的电邮:")
		var email = ""
		fmt.Scanln(&email)

		newCustomer := model.NewCustomer1(name, gender, age, phone, email)
		this.customerService.UpdateCustomer(newCustomer, res)
	}
}

// 退出管理系统
func (this *customerView) exit() {
	fmt.Println("退出 请输入y, 取消退出请输入n 其他键无效。")
	var key = ""
	for {
		fmt.Scanln(&key)
		if key == "y" || key == "Y" {
			this.loop = false
			break
		} else if key == "n" || key == "N" {
			this.loop = true
			fmt.Println("您 取消退出成功， 请继续操作。")
			break
		} else {
			fmt.Println("您的 输入有误，请重新输入 y or n ")
		}
	}
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
			this.add()
		case "2":
			this.update()
		case "3":
			fmt.Println("删除客户")
			this.delete()
		case "4":
			fmt.Println("客户列表")
			this.list()
		case "5":
			this.exit()
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
	cv := &customerView{
		"",
		true,
		service.NewCustomerService(),
	}

	cv.mainMenu()
}
