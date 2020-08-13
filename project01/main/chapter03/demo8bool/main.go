package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1. bool类型大小占用1字节空间
	//2. bool类型只能取true or false
	var b = false
	fmt.Println("b=", b)
	fmt.Println("b占用的空间 = ", unsafe.Sizeof(b))
}
