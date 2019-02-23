package main

import "fmt"

// 定义一个结构体Account
type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}


// 方法 存款
func (a *Account) Deposite(money float64, pwd string){
	// 查看下输入的密码是不是正确。
	if a.Pwd != pwd {
		fmt.Println("你 输入的密码不正确")
		return
	}

	// 看看存款金额是否正确
	if money <= 0 {
		fmt.Println("您 输入的金额不正确。")
		return
	}
	a.Balance += money
	fmt.Println("存款成功。")
}

func (a *Account) WithDraw(money float64, pwd string){
	if a.Pwd != pwd {
		fmt.Println("您输入的密码有误.")
		return
	}

	if money > a.Balance {
		fmt.Println("您 输入的金额不正确。")
		return
	}
	a.Balance -= money
	fmt.Println("取款成功。")
}

func (a *Account) Query(pwd string){
	// 看下输入的密码是不是正确。
	if a.Pwd != pwd {
		fmt.Println("您输入的密码不正确。")
		return
	}

	fmt.Println("您的余额为:",a.Balance)
}

func main(){
	// 测试
	account := &Account{
		AccountNo : "abc",
		Pwd : "666666",
		Balance : 100.0,
	}

	// account.Query("6661666")
	account.Query("666666")
	
	account.Deposite(99,"666666")
	account.Query("666666")
	
	account.WithDraw(10,"666666")
	account.Query("666666")
}