package model

import (
	"fmt"
)

//声明一个Customer结构体，表示一个客户信息

//Customer 客户结构体
type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//NewCustomer 编写一个工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//NewCustomer2 不需要传递ID的工厂模式
func NewCustomer2(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//GetInfo 返回用户的信息
func (customer Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", customer.ID, customer.Name, customer.Gender,
		customer.Age, customer.Phone, customer.Email)
	return info
}
