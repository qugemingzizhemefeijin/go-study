package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
字符串是一个不可改变的字节序列，字符串底层可以看作是一个只读的二进制数组。
 字符串底层结构体
 type StringHeader struct {
	Data uintptr
	Len int
 }
*/
func main() {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println(s, hello, s1, world, s2)

	// 字符串和数组类似，内置的len函数返回字符串的长度
	fmt.Println(len(s))

	// 也可以通过reflect.StringHeader结构访问字符串的长度（不推荐）
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len)
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len)

	// 可以通过将类型转为字节类型查看字符串底层对应的数据
	fmt.Println("%#v\n", []byte("Hello, 世界"))

	// 如果不想解码UTF8字符串，想直接遍历原始的字节码，可以将字符串强制转为[]byte字节序列后再进行遍历
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

	fmt.Println("===========================")

	// 还可以将字符串和[]rune类型相互转换
	//[]rune其实是[]int32类型，这里的rune只是int32类型的别名，并不是重新定义的类型。rune用于标识每个Unicode码点，目前只使用了21个bit位。
	fmt.Printf("%#v\n", []rune("世界"))             // []int32{19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // "世界"

}
