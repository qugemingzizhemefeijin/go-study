package main

// 必须添加下面的，否则会报 could not determine kind of name for C.puts

//#include <stdio.h>
import "C"

// Go语⾔通过⾃带的⼀个叫CGO的⼯具来⽀持C语⾔函数调⽤， 同时我们可以⽤Go语⾔导出C动态库接⼝给其它语⾔使⽤。
// 基于C标准库函数输出字符串
func main() {
	// 代码通过 import "C" 语句启⽤CGO特性， 主函数只是通过Go内置的println函数输出字符串， 其中并没有任何和CGO相关的代码。 虽然没有调⽤CGO的相关函数，
	// 但是go build命令会在编译和链接阶段启动gcc编译器，这已经是⼀个完整的CGO程序了
	println("hello cgo")

	// 基于C标准库函数输出字符串
	C.puts(C.CString("Hello, World\n"))

	// 我们不仅仅通过 import "C" 语句启⽤CGO特性， 同时包含C语⾔的 <stdio.h> 头⽂件。 然后通过CGO包的 C.CString 函数将Go语⾔字符串转为C语⾔字符串，
	// 最后调⽤CGO包的 C.puts 函数向标准输出窗⼝打印转换后的C字符串。
	// 我们没有在程序退出前释放 C.CString 创建的C语⾔字符串，会导致内存泄漏。
	// 但是对于这个⼩程序来说， 这样是没有问题的， 因为程序退出后操作系统会⾃动回收程序的所有资源。
}