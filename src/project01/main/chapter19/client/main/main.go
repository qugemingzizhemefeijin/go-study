package main

import (
	"fmt"
	"go_code/project01/main/chapter19/client/processes"
	"os"
)

//定义两个变量，一个表示用户ID，一个表示用户密码，一个表示用户昵称
var userID int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true
	for {
		fmt.Println("-------------------------------欢迎登陆多人聊天系统-------------------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			//fmt.Println("登陆聊天室")
			//loop = false
			fmt.Println("请输入用户的ID:")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			// 完成登录
			//1. 创建一个UserProcess的实例
			up := &processes.UserProcess{}
			up.Login(userID, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字(nickname):")
			fmt.Scanf("%s\n", &userName)

			//2. 调用UserProcee，完成注册的请求
			up := &processes.UserProcess{}
			up.Register(userID, userPwd, userName)
		case 3:
			//loop = false
			fmt.Println("客户端退出系统...")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	//根据用户的输入显示新的提示信息
	// if key == 1 {
	// 	//说明用户要登陆
	// 	fmt.Println("请输入用户的ID:")
	// 	fmt.Scanf("%d\n", &userID)
	// 	fmt.Println("请输入用户的密码:")
	// 	fmt.Scanf("%s\n", &userPwd)
	// 	//先把登录的函数，写到另外一个文件
	// 	//此处会报错，需要到 GOPATH目录下执行  go build -o client.exe go_code/project01/main/chapter19/client/main 或者 go run *.go
	// 	// err := Login(userID, userPwd)
	// 	// if err != nil {
	// 	// 	fmt.Println("登录失败")
	// 	// } else {
	// 	// 	fmt.Println("登录成功")
	// 	// }
	// 	//这里我们会需要重新调用
	// 	//Login(userID, userPwd)
	// } else if key == 2 {
	// 	fmt.Println("进行用户注册的逻辑...")
	// }
}
