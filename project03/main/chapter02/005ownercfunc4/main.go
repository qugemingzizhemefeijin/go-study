package main

// 在Go1.10中CGO新增加了⼀个 _GoString_ 预定义的C语⾔类型， ⽤来表示Go语⾔字符串

//void SayHello(_GoString_ s);
import "C"

import "fmt"

func main() {
	// 虽然看起来全部是Go语⾔代码， 但是执⾏的时候是先从Go语⾔的 main 函数， 到CGO⾃动⽣成的C语⾔版本 SayHello 桥接函数，
	// 最后⼜回到了Go语⾔环境的 SayHello 函数。 这个代码包含了CGO编程的精华， 读者需要深⼊理解。
	C.SayHello("Hello, World\n")
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}