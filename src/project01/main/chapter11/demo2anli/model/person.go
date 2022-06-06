package model

import "fmt"

type person struct {
	Name   string
	age    int
	salary float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age 和 salary ，我们编写一对set的方法和get的方法

func (p *person) SetAge(age int) {
	if age <= 0 || age > 150 {
		fmt.Println("年龄不合理")
	} else {
		p.age = age
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	if salary < 3000 || salary > 30000 {
		fmt.Println("薪水设置错误")
	} else {
		p.salary = salary
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}
