package main

import (
	"fmt"
	Tools "go_code/project01/main/chapter06/demo1/utils" //前面Tools相当于是取了个别名，utils则不可使用了
)

// import时，路径从$GOPATH的src下开始，不用带src，编译器会自动从src下开始引入；
// 在同一个包下，不能有相同的函数名，否则报重复定义；
// 如果你要编译成一个可执行程序文件，就需要将这个包声明为main，即package main。

// 在GOPATH中 编译 go build go_code/project/main [main包存在的路径下]
// go build -o bin/my.exe go_code/project/main
func main() {
	var n1, n2 float64 = 1.2, 2.3
	result := Tools.Calc(n1, n2, '+')
	fmt.Println("res=", result)
}
