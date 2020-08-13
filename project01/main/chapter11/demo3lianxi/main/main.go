package main

import (
	"fmt"
	"go_code/project01/main/chapter11/demo3lianxi/model"
)

func main() {
	account := model.NewAccount("gs12341121256")
	if account == nil {
		fmt.Println("创建帐号失败")
	} else {
		fmt.Println(account)
		account.SetPwd("12345")
		account.SetBalance(2002)
		account.SetPwd("abcdef")
		fmt.Printf("accountNo:%v,balance=%v,pwd=%v\n", account.GetNo(), account.GetBalance(), account.GetPwd())
	}
}
