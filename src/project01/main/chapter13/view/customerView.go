package main

import (
	"fmt"
	"project01/main/chapter13/model"
	"project01/main/chapter13/service"
)

type customerView struct {
	//定义必要的字段
	key  string //接收用户输入...
	loop bool   //表示是否循环的显示主菜单
	//增加一个字段customerService
	customerService *service.CustomerService
}

//list 显示所有的客户信息
func (cv *customerView) list() {
	//首先获取到当前所有的客户信息
	customers := cv.customerService.List()
	//显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].ID, "\t", customers[i].Name)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

//add 得到用户的输入信息构建新的客户，并完成添加
func (cv *customerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意：ID号没有让用户输入，ID是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)

	//调用service add方法
	if cv.customerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

//delete 得到用户的输入ID，删除该ID对应的客户
func (cv *customerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}

	fmt.Println("确认是否删除(Y/N)：")
	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
	for {
		choice := ""
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			//调用customerService 的 Delete方法
			if cv.customerService.Delete(id) {
				fmt.Println("---------------------删除完成---------------------")
			} else {
				fmt.Println("---------------------删除失败，输入的id号不存在----")
			}
			break
		} else if choice == "n" || choice == "N" {
			break
		} else {
			fmt.Println("确认是否删除(Y/N)：")
		}
	}
}

func (cv *customerView) exit() {
	fmt.Println("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "Y" || cv.key == "y" || cv.key == "N" || cv.key == "n" {
			break
		}

		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
	}

	if cv.key == "Y" || cv.key == "y" {
		cv.loop = false
	}
}

//update 修改
func (cv *customerView) update() {
	fmt.Println("---------------------修改客户---------------------")
	fmt.Println("请选择待修改客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃修改操作
	}

	customer := cv.customerService.FindByCustomer(id)
	if customer == nil {
		fmt.Println("你输入的客户编号不存在")
		return
	}

	//如果编号存在，则执行修改操作
	fmt.Println("姓名(", customer.Name, "):")
	name := ""
	fmt.Scanln(&name)
	if name == "" {
		name = customer.Name
	}
	fmt.Println("性别(", customer.Gender, "):")
	gender := ""
	if gender == "" {
		gender = customer.Gender
	}
	fmt.Scanln(&gender)
	fmt.Println("年龄(", customer.Age, "):")
	age := 0
	fmt.Scanln(&age)
	if age == 0 {
		age = customer.Age
	}
	fmt.Println("电话(", customer.Phone, "):")
	phone := ""
	fmt.Scanln(&phone)
	if phone == "" {
		phone = customer.Phone
	}
	fmt.Println("邮箱(", customer.Email, "):")
	email := ""
	fmt.Scanln(&email)
	if email == "" {
		email = customer.Email
	}

	//构建一个新的Customer实例
	//注意：ID号没有让用户输入，ID是唯一的，需要系统分配
	newCustomer := model.NewCustomer(customer.ID, name, gender, age, phone, email)

	//调用Service的Update方法
	//调用service add方法
	if cv.customerService.Update(newCustomer) {
		fmt.Println("---------------------修改完成---------------------")
	} else {
		fmt.Println("---------------------修改失败---------------------")
	}
}

//mainMenu 显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !cv.loop {
			break
		}
	}

	fmt.Println("你退出了客户关系管理系统...")
}

//要求：
//实现对客户对象的插入、修改和删除（用切片实现），并能够打印客户明细表
func main() {
	//在主函数中创建一个customerView实例，并运行显示主菜单
	var cv customerView = customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}
	cv.mainMenu()
}
