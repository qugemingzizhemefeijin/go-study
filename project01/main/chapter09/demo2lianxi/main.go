package main

import (
	"fmt"
)

func modifyUser(users map[string]map[string]string, name string) {
	user, ok := users[name]
	//也可以用 if users[name] != nil //来判断
	if ok {
		fmt.Printf("用户名 %v 存在，修改密码为888888\n", name)
		user["pwd"] = "888888"
	} else {
		fmt.Printf("用户名 %v 不存在，新增用户\n", name)
		var nickname, pwd string
		fmt.Println("请输入你的昵称：")
		fmt.Scanln(&nickname)
		fmt.Println("请输入你的密码：")
		fmt.Scanln(&pwd)

		user := map[string]string{"nickname": nickname, "pwd": pwd}
		users[name] = user
	}
}

//课堂练习
func main() {
	//1) 使用  map[string]map[string]string的map类型
	//2) key:表示用户名 是唯一的，不可以重复
	//3) 如果某个用户名存在，就将其密码修改"888888"，如果不存在就增加这个用户信息（包括昵称nickname 和密码pwd）
	//4) 编写一个函数 modifyUser(users map[string]map[string]string, name string)完成上述功能

	//1.首先定义一个结构体
	var users = make(map[string]map[string]string, 10)
	fmt.Println(users)
	modifyUser(users, "kitty")
	modifyUser(users, "kitty")
	modifyUser(users, "amy")
	fmt.Println(users)
}
