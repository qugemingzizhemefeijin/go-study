package service

import (
	"project01/main/chapter13/model"
)

//CustomerService 该结构体完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的ID+1
	customerNum int
}

//NewCustomerService 编写一个方法，可以返回 *CustomerService类型
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1

	customer := model.NewCustomer(1, "张三", "男", 20, "119", "zhangsan@sohu.com")
	customerService.customers = append(customerService.customers, customer)

	return customerService
}

//List 返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

//Add 添加客户到切片中
func (cs *CustomerService) Add(customer model.Customer) bool {
	//我们确定一个分配ID的规则，就是添加的顺序
	cs.customerNum++
	customer.ID = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

//FindByID 根据ID查找客户在切片中对应下标，如果没有该客户，返回-1
func (cs *CustomerService) FindByID(id int) int {
	//遍历customers切片
	for index, customer := range cs.customers {
		if customer.ID == id {
			return index
		}
	}

	return -1
}

//FindByCustomer 根据ID查找Customer实例，如果没有则返回nil
func (cs *CustomerService) FindByCustomer(id int) *model.Customer {
	//遍历customers切片
	for _, customer := range cs.customers {
		if customer.ID == id {
			return &customer
		}
	}

	return nil
}

//Delete 根据ID删除客户(从切片中删除)
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindByID(id)
	//如果index == -1，说明没有这个客户，直接返回false
	if index == -1 {
		return false
	}
	//如何从切片中你删除一个元素呢？
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

//Update 修改客户信息（修改切片中信息）
func (cs *CustomerService) Update(customer model.Customer) bool {
	index := cs.FindByID(customer.ID)
	//如果index == -1，说明没有这个客户，直接返回false
	if index == -1 {
		return false
	}

	cs.customers[index] = customer
	return true
}
