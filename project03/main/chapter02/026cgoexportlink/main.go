package main

// CGO不仅可以使⽤C静态库， 也可以将Go实现的函数导出为C静态库。 我们现在⽤Go实现前⾯的number库的模加法函数。

import "C"

// CGO不仅可以使⽤C静态库， 也可以将Go实现的函数导出为C静态库。 我们现在⽤Go实现前⾯的number库的模加法函数。采⽤以下命令构建：
// go build -buildmode=c-archive -o number.a
// 在⽣成number.a静态库的同时， cgo还会⽣成⼀个number.h⽂件。

// 再编写_test_main.c
// gcc -o a.out _test_main.c number.a
// 执行编译后的文件

// CGO导出动态库的过程和静态库类似， 只是将构建模式改为 c-shared ， 输出⽂件名改为 number.so ⽽已：
// go build -buildmode=c-shared -o number.so
// gcc -o a2.out _test_main.c number.so
func main() {

}

//export number_add_mod
func number_add_mod(a, b, mod C.int) C.int {
	return (a + b) % mod
}
