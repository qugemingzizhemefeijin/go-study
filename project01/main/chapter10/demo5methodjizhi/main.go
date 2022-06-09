package main

import (
	"fmt"
	"math"
)

//编写一个结构体Circle，字段为radius，声明一个area和Circle绑定，可以返回面积

//Circle  ...
type Circle struct {
	Radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func main() {
	circel := Circle{1.0}
	fmt.Printf("圆的面积为：%.2f\n", circel.area())
}
