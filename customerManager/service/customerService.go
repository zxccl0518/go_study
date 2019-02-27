package service

import (
	"fmt"

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

func (this *CustomerService) List() {
	for _, v := range this.customers {
		fmt.Println(v)
	}
}
