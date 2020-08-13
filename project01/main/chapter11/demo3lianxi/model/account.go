package model

import (
	"fmt"
	"unicode/utf8"
)

type account struct {
	accountNo string
	balance   float64
	pwd       string
}

//NewAccount 工厂模式
func NewAccount(accountNo string) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("帐号的长度不对...")
		return nil
	}

	return &account{
		accountNo: accountNo,
		pwd:       "666666",
		balance:   0.0,
	}
}

func (a *account) SetBalance(balance float64) {
	if balance <= 20 {
		fmt.Println("余额设置有误，必须要大于20")
	} else {
		a.balance = balance
	}
}

func (a *account) GetBalance() float64 {
	return a.balance
}

func (a *account) SetPwd(pwd string) {
	if utf8.RuneCountInString(pwd) != 6 {
		fmt.Println("密码设置有误，必须要设置6位密码")
	} else {
		a.pwd = pwd
	}
}

func (a *account) GetPwd() string {
	return a.pwd
}

func (a *account) GetNo() string {
	return a.accountNo
}
