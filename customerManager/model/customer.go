package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂模式 创建customer实例
func NewCustomer(id int, name string, sex string, age int, phone string, email string) Customer {
	return Customer{
		id,
		name,
		sex,
		age,
		phone,
		email,
	}
}

// 工厂模式 创建customer实例
func NewCustomer1(name string, sex string, age int, phone string, email string) Customer {
	return Customer{
		2,
		name,
		sex,
		age,
		phone,
		email,
	}
}

func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Phone)
	return info
}
