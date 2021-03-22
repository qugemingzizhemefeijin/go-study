package main

import (
	"fmt"
	"project01/main/chapter12/utils"
)

func main() {
	fmt.Println("这个是面向对象的方式完成的~~~")
	account := utils.NewFamilyAccount()
	account.MainMenu()
}
