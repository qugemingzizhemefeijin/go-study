package main

import (
	"fmt"
)

//获取用户输入的数据
//案例 要求：从控制台接收用户系你想【姓名、年龄、薪水、是否通过考试】
func main() {
	//1.使用fmt.Scanln()获取
	var name string
	var age byte
	var salary float32
	var pass bool

	// fmt.Println("请输入姓名 ")
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄 ")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水 ")
	// fmt.Scanln(&salary)

	// fmt.Println("请输入是否通过 ")
	// fmt.Scanln(&pass)

	// fmt.Printf("姓名\t年龄\t薪水\t是否通过\n")
	// fmt.Printf("%v\t%d\t%v\t%t\t\n", name, age, salary, pass)
	//2.使用fmt.Scanf()获取
	fmt.Println("请输入你的姓名、年龄、薪水、是否通过考试")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &pass)

	fmt.Printf("姓名\t年龄\t薪水\t是否通过\n")
	fmt.Printf("%v\t%d\t%v\t%t\t\n", name, age, salary, pass)
}
