package service

import (
	"github.com/zxccl0518/go_study/customerManager/model"
)

// 该 接头体完成对customer操作，包括增删改查。
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	// 该字段后面，还可以作为新客户的id+1
	customerNum int
}

func NewCustomerService() *CustomerService {
	customService := CustomerService{}
	// 为了能够看到客户在切片中，我们初始化一个客户。
	customService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@souhu.com")
	customService.customers = append(customService.customers, customer)

	return &customService
}

// 返回客户列表的一个切片
func (this *CustomerService) List() []model.Customer {
	// for _, v := range this.customers {
	// 	fmt.Println(v)
	// }
	return this.customers
}

// 增加一个新用户
func (this *CustomerService) AddCustomer(cus model.Customer) {
	this.customerNum++
	cus.Id = this.customerNum
	this.customers = append(this.customers, cus)
}

// 删除客户
func (this *CustomerService) DeleteCustomer(id int) {
	this.customers = append(this.customers[:id], this.customers[id+1:]...)
}

// 修改客户信息
func (this *CustomerService) UpdateCustomer(cus model.Customer, index int) {
	this.customers[index].Name = cus.Name
	this.customers[index].Gender = cus.Gender
	this.customers[index].Age = cus.Age
	this.customers[index].Phone = cus.Phone
	this.customers[index].Email = cus.Email
}

// 根据id 查找客户
func (this *CustomerService) FindById(id int) int {
	for k, v := range this.customers {
		if v.Id == id {
			return k
		}
	}
	return -1
}
