package main

import (
	"fmt" //如果我们没有使用到一个包，但是又不想去掉，前面加一个 _，表示忽略
)

//演示golang中基本数据类型的转换
func main() {
	//GO中数据类型的转换必须显示转换，如 float32(v), uint8(v)等
	var a uint8 = 10
	//希望将 i => float类型
	var n1 float32 = float32(a)
	var n2 int16 = int16(a)

	var i int8 = -1
	//n3溢出造成符号位丢失
	var n3 uint16 = uint16(i)

	fmt.Printf("a=%v, n1=%v, n2=%v, n3=%v\n", a, n1, n2, n3)

	//
	var n10 int32 = 12
	var n11 int8
	//var n12 int8
	n11 = int8(n10) + 127
	//n12 = int8(n10) + 128 此处编译不通过 128超出了int8范围
	fmt.Println(n11)
}
