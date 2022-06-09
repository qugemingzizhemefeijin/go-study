package main

// 使用C静态库
// gcc -c -o number.o number.c
// ar rcs libnumber.a number.o
// ⽣成libnumber.a静态库之后， 我们就可以在CGO中使⽤该资源了。

/*
其中有两个#cgo命令， 分别是编译和链接参数。
CFLAGS通过 -I./number 将number库对应头⽂件所在的⽬录加⼊头⽂件检索路径。
LDFLAGS通过 -L${SRCDIR}/number 将编译后number静态库所在⽬录加为链接库检索路径， -lnumber 表示链接libnumber.a静态库。

需要注意的是， 在链接部分的检索路径不能使⽤相对路径（C/C++代码的链接程序所限制） ，
我们必须通过cgo特有的 ${SRCDIR} 变量将源⽂件对应的当前⽬录路径展开为绝对路径（因此在windows平台中绝对路径不能有空⽩符号）。

因为我们有number库的全部代码， 所以我们可以⽤go generate⼯具来⽣成静态库， 或者是通过Makefile来构建静态库。
因此发布CGO源码包时， 我们并不需要提前构建C静态库。

因为多了⼀个静态库的构建步骤， 这种使⽤了⾃定义静态库并已经包含了静态库全部代码的Go包⽆法直接⽤go get安装。
不过我们依然可以通过go get下载， 然后⽤go generate触发静态库构建， 最后才是go install来完成安装。

为了⽀持go get命令直接下载并安装， 我们C语⾔的 #include 语法可以将number库的源⽂件链接到当前的包。

创建 z_link_number_c.c ⽂件如下:
#include "./number/number.c"

然后在执⾏go get或go build之类命令的时候， CGO就是⾃动构建number库对应的代码。
这种技术是在不改变静态库源代码组织结构的前提下， 将静态库转化为了源代码⽅式引⽤。 这种CGO包是最完美的。
*/

//#cgo CFLAGS: -I./number
//#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
//
//#include "number.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.number_add_mod(10, 5, 12))
}
