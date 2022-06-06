package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0' // 字符'0'

	//当我们直接输出byte值，就是输出了对应的字符的码值
	// 'a' => 97
	// '0' => 48
	// 'A' => 63
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果我们希望输出对应的字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c \n", c1, c2)

	//var c3 byte = '北' //overflow溢出错误
	var c3 int = '北' //就可以了
	fmt.Printf("c3=%c c3对应码值=%d \n", c3, c3)

	//字符是以字节存储的
	//字符使用UTF-8编码
}
