package utils

import (
	"fmt"
)

//Age 11
var Age int

//Name 22
var Name string

//这里给初始化
func init() {
	fmt.Println("utils init()")
	Age = 20
	Name = "tom"
}
