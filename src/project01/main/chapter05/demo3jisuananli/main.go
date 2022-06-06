package main

import (
	"fmt"
)

func main() {
	//参加百米运动会，如果用时8秒以内进入决赛，否则提示淘汰。并且根据性别提示进入男子组或女子组。输入成绩和性别，进行判断。
	// var second float64
	// var gender string

	// fmt.Println("请输入性别(男/女)：")
	// fmt.Scanln(&gender)
	// fmt.Println("情输入成绩(秒)：")
	// fmt.Scanln(&second)

	// if second <= 8 {
	// 	if gender == "男" {
	// 		fmt.Println("你进入男子组决赛")
	// 	} else {
	// 		fmt.Println("你进入女子组决赛")
	// 	}
	// } else {
	// 	fmt.Println("您已经被淘汰了")
	// }

	//出票系统：根据淡旺季的月份和年龄，打印票价
	// 4-10旺季：成人(18-60):60，儿童(<18):半价,老人>60:1/3
	// 淡季：成人40，其他：20
	var month uint8
	var age uint8

	fmt.Println("请输入月份：")
	fmt.Scanln(&month)

	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	var price uint8
	if month >= 4 && month <= 10 {
		if age >= 18 && age <= 60 {
			price = 60
		} else if age < 18 {
			price = 30
		} else {
			price = 20
		}
	} else {
		if age >= 18 && age <= 60 {
			price = 40
		} else {
			price = 20
		}
	}

	fmt.Printf("当前你需要支付的票价为：%d\n", price)
}
