package main

/*
#include<stdio.h>

void printString(const char* s, int n) {
	int i;
	for(i = 0; i < n; i++) {
		putchar(s[i]);
	}
	putchar('\n');
}
 */
import "C"
import (
	"reflect"
	"unsafe"
)

func printString(s string) {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	C.printString((*C.char)(unsafe.Pointer(p.Data)), C.int(len(s)))

	// 错误的代码
	//tmp := uintptr(unsafe.Pointer(&x))
	//pb := (*int16)(unsafe.Pointer(tmp))
	//*pb = 42

	// 因为 tmp 并不是指针类型，在它获取到 Go 对象地址之后 x 对象可能会被移动，但是因为不是指针类型，所以不会被 Go 语言运行时更新成新内存的地址。
	// 在非指针类型的 tmp 保持 Go 对象的地址，和在 C 语言环境保持 Go 对象的地址的效果是一样的：
	// 如果原始的 Go 对象内存发生了移动，Go 语言运行时并不会同步更新它们。
}

// 为了简化并高效处理上一个例子向 C 语言传入 Go 语言内存的问题，cgo 针对该场景定义了专门的规则：
// 在 CGO 调用的 C 语言函数返回前，cgo 保证传入的 Go 语言内存在此期间不会发生移动，C 语言函数可以大胆地使用 Go 语言的内存！

// CGO 的这种看似完美的规则也是存在隐患的。我们假设调用的 C 语言函数需要长时间运行，那么将会导致被他引用的 Go 语言内存在 C 语言返回前不能被移动，
// 从而可能间接地导致这个 Go 内存栈对应的 goroutine 不能动态伸缩栈内存，也就是可能导致这个 goroutine 被阻塞。
// 因此，在需要长时间运行的 C 语言函数（特别是在纯 CPU 运算之外，还可能因为需要等待其它的资源而需要不确定时间才能完成的函数），需要谨慎处理传入的 Go 语言内存。

// 不过需要小心的是在取得 Go 内存后需要马上传入 C 语言函数，不能保存到临时变量后再间接传入 C 语言函数。
// 因为 CGO 只能保证在 C 函数调用之后被传入的 Go 语言内存不会发生移动，它并不能保证在传入 C 函数之前内存不发生变化。
func main() {
	s := "hello, world"
	printString(s)
}
