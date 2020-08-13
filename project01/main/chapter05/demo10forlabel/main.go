package main

import (
	"fmt"
)

func main() {
	//编写一个无限循环的空值，然后不停的随机生成数，当生成了99时，就退出这个无限循环
	// fmt.Println(time.Now())
	// rand.Seed(time.Now().Unix())
	// n := rand.Intn(100) + 1
	// fmt.Println(n)

	// var count int = 0
	// for {
	// 	rand.Seed(time.Now().UnixNano())
	// 	n := rand.Intn(100) + 1
	// 	count++
	// 	if n == 99 {
	// 		break
	// 	}
	// }
	// fmt.Println("生成 99 一共使用了 ", count)

	//这里演示一下指定标签的形式来使用break
	// label2: //设置一个标签
	// 	for i := 0; i < 4; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 2 {
	// 				break label2
	// 			}
	// 			fmt.Println("j=", j)
	// 		}
	// 	}

	//1.100以内的数求和，求出当和第一次大于20的当前数
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum >= 20 {
			fmt.Printf("和第一次大于20的当前数是：%d\n", i)
			break
		}
	}

	//2.实现登录验证，有三次机会，如果用户名为：张无忌，密码 888 提示登录成功，否则提示还有几次机会
	for i := 1; i <= 3; i++ {
		var username, password string
		fmt.Println("请输入用户名:")
		fmt.Scanln(&username)
		fmt.Println("请输入密码:")
		fmt.Scanln(&password)
		if username == "张无忌" && password == "888" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Printf("用户名或密码输入错误，你还有%d次机会\n", 3-i)
		}
	}
}
