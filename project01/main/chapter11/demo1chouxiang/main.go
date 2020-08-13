package main

import (
	"fmt"
)

//Account 定义一个结构体Account
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//方法

//SaveMoney 1.存款
func (account *Account) SaveMoney(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功，余额为：", account.Balance)
}

//WithDraw 2.取款
func (account *Account) WithDraw(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看存款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功，余额为：", account.Balance)
}

//Query 3.查询余额
func (account *Account) Query(pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的帐号为=%v 余额为%v\n", account.AccountNo, account.Balance)
}

func main() {
	//测试一把
	account := Account{
		AccountNo: "gs111111",
		Pwd:       "666666",
		Balance:   100.0,
	}

	//这里可以做得更灵活，让用户通过控制台输入命令来进行操作...
	account.Query("666666")
	account.SaveMoney(200, "666666")
	account.Query("666666")
	account.WithDraw(155, "666666")
	account.Query("666666")
}
