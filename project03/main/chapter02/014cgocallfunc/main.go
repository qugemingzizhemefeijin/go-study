package main

/*
#include <errno.h>

static int add(int a, int b) {
	return a + b;
}

// CGO也针对 <errno.h> 标准库的 errno 宏做的特殊⽀持： 在CGO调⽤C函数时如果有两个返回值， 那么第⼆个返回值将对应 errno 错误状态。
// 我们可以近似地将div函数看作为以下类型的函数：func C.div(a, b C.int) (C.int, [error])
// 第⼆个返回值是可忽略的error接⼝类型， 底层对应 syscall.Errno 错误类型。
static int div(int a, int b) {
	if (b == 0) {
		errno = EINVAL;
		return 0;
	}
	return a / b;
}

// C 语言函数还有一种没有返回值类型的函数，用 void 表示返回值类型。一般情况下，我们无法获取 void 类型函数的返回值，因为没有返回值可以获取。
// 前面的例子中提到，cgo 对 errno 做了特殊处理，可以通过第二个返回值来获取 C 语言的错误状态。对于 void 类型函数，这个特性依然有效。
static void noreturn() {

}
 */
import "C"
import "fmt"

func main() {
	fmt.Println(C.add(1, 1))

	v0, err0 := C.div(6, 3)
	fmt.Println(v0, err0)

	v1, err1 := C.div(2, 0)
	fmt.Println(v1, err1) // 这里按理应该返回“invalid argument”，不知道为什么返回的是“The device does not recognize the command”

	_, err := C.noreturn()
	fmt.Println(err)

	// 我们也可以尝试获取第一个返回值，它对应的是 C 语言的 void 对应的 Go 语言类型：
	v, _ := C.noreturn()
	fmt.Printf("%#v\n", v) // 输出 main._Ctype_void{}

	// 我们可以看出 C 语言的 void 类型对应的是当前的 main 包中的 _Ctype_void 类型。
	// 其实也将 C 语言的 noreturn 函数看作是返回 _Ctype_void 类型的函数，这样就可以直接获取 void 类型函数的返回值：
	fmt.Println(C.noreturn()) // 输出 [] _Ctype_void 类型对应一个 0 长的数组类型 [0]byte，因此 fmt.Println 输出的是一个表示空数值的方括弧。
}
