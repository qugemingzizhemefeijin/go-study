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
	fmt.Println("电话开始")
}

func (p Phone) Stop() {
	fmt.Println("电话结束")
}

type Camcra struct {
	Name string
}

func (c Camcra) Start() {
	fmt.Println("相机开始")
}

func (c Camcra) Stop() {
	fmt.Println("相机结束")
}

//多态数组
func main() {
	var usbArr [3]Usb

	usbArr[0] = Phone{"华为"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camcra{"佳能"}

	fmt.Println(usbArr)
}
