package main

import (
	"fmt"
)

func main() {
	//string的基本使用
	var address string = "北京长城 Hello World"
	fmt.Println("address = ", address)

	//1.字符串的字节使用UTF-8编码标识Unicode文本。
	//2.字符串一旦赋值，则是不可变的
	//3.字符串表示方式有两种，1双引号,2反引号
	var name string = `我是你的'"var name = string朋友`
	fmt.Println("name = ", name)

	sex := `你是男的`
	fmt.Println("sex = ", sex)

	//可以直接输出
	str3 := `package main

	import (
		"fmt"
	)
	
	func main() {
		//string的基本使用
		var address string = "北京长城 Hello World"
		fmt.Println("address = ", address)
	}
	`
	fmt.Println("str3 = ", str3)

	//4.字符串拼接方式
	var p1 string = "Hello " + "World"
	p1 += " haha"
	fmt.Println(p1)

	//5.当一个拼接的操作很长时，可以分行写。+号必须要留在上面
	var str2 = "Hello," +
		"World"
	fmt.Println(str2)
}
