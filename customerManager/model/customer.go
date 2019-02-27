package model

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
