package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("你好，GOLANG")
	}

	//for循环的第二种写法
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	//第三种写法
	// for {
	// 	//死循环
	// }

	//字符串遍历方式1-传统方式
	var str string = "hello,world!北京"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	//字符串遍历方式，解决中文问题
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	//字符串遍历方式2-for-range
	for index, val := range str {
		fmt.Printf("index=%d,val=%c \n", index, val)
	}
}
