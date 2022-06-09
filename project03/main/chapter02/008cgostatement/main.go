package main

// 其中在windows平台下， 编译前会预定义X86宏为1； 在⾮widnows平台下， 在链接阶段会要求链接
// math数学库。 这种⽤法对于在不同平台下只有少数编译选项差异的场景⽐较适⽤。
// 如果在不同的系统下cgo对应着不同的c代码， 我们可以先使⽤ #cgo 指令定义不同的C语⾔的宏， 然后
// 通过宏来区分不同的代码：

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#if defined(CGO_OS_WINDOWS)
const char* os = "windows";
#elif defined(CGO_OS_DARWIN)
static const char* os = "darwin";
#elif defined(CGO_OS_LINUX)
static const char* os = "linux";
#else
# error(unknown os)
#endif
 */
import "C"

func main() {
	print(C.GoString(C.os))
}
