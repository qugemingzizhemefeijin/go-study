package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "电话开始")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "电话结束")
}

func (p Phone) Call() {
	fmt.Println(p.Name, "电话拨打")
}

type Camcra struct {
	Name string
}

func (c Camcra) Start() {
	fmt.Println(c.Name, "相机开始")
}

func (c Camcra) Stop() {
	fmt.Println(c.Name, "相机结束")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//如果usb是指向Phone结构体变量，则还需要调用Call方法
	//类型断言来咯。。。
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
	fmt.Println()
}

//使用类型断言，在调用Phone的时候，多执行一个Call代码
func main() {
	var usbArr [3]Usb

	usbArr[0] = Phone{"华为"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camcra{"佳能"}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
