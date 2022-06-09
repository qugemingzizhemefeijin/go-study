package main

import (
	"fmt"
)

//声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

//让Phone实现USB接口的方法
func (phone Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (phone Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
}

//让Camera 实现Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

//计算机
type Computer struct {
}

//编写一个方法Working，此方法接收一个Usb接口类型变量
func (c *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	//测试
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}
