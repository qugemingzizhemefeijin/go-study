package main

import (
	"fmt"
	"go_code/project01/main/chapter11/demo2anli/model"
)

func main() {
	//设计一个person.go，不能随便查看人的年龄，工资等隐私，并对输入的年龄进行合理的验证
	var p = model.NewPerson("小李")
	p.SetAge(25)
	p.SetSalary(12000)

	fmt.Println(p)
	fmt.Printf("%v 的 年龄是： %v岁\n", p.Name, p.GetAge())
	fmt.Printf("%v 的 薪水是： %v元\n", p.Name, p.GetSalary())
}
