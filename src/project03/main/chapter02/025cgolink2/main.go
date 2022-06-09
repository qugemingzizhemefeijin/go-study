package main

// 使⽤C动态库
/*
动态库出现的初衷是对于相同的库， 多个进程可以共享同⼀个， 以节省内存和磁盘资源。
但是在磁盘和内存已经⽩菜价的今天， 这两个作⽤已经显得微不⾜道了， 那么除此之外动态库还有哪些存在的价值呢？
从库开发⻆度来说， 动态库可以隔离不同动态库之间的关系， 减少链接时出现符号冲突的⻛险。
⽽且对于windows等平台， 动态库是跨越VC和GCC不同编译器平台的唯⼀的可⾏⽅式。

对于CGO来说，使⽤动态库和静态库是⼀样的，因为动态库也必须要有⼀个⼩的静态导出库⽤于链接动态库
（Linux下可以直接链接so⽂件，但是在Windows下必须为dll创建⼀个 .a ⽂件⽤于链接）。

gcc -shared -o libnumber.so number.c
 */

//#cgo CFLAGS: -I./number
//#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
//
//#include "number.h"
import "C"
import "fmt"

func main() {
	// 编译时GCC会⾃动找到libnumber.a或libnumber.so进⾏链接
	fmt.Println(C.number_add_mod(10, 5, 12))
}
