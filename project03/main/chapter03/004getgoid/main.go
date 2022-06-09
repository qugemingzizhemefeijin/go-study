package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// Go语⾔刻意没有提供goid的原因是为了避免被滥⽤。 因为⼤部分⽤户在轻松拿到goid之后， 在之后的编程中会不⾃觉地编写出强依赖goid的代码。
// 强依赖goid将导致这些代码不好移植， 同时也会导致并发模型复杂化。 同时， Go语⾔中可能同时存在海量的Goroutine，
// 但是每个Goroutine何时被销毁并不好实时监控， 这也会导致依赖goid的资源⽆法很好地⾃动回收（需要⼿⼯回收） 。
// 不过如果你是Go汇编语⾔⽤户， 则完全可以忽略这些借⼝。

// 纯Go⽅式获取goid
// 使⽤纯Go的⽅式获取goid的⽅式虽然性能较低，但是代码有着很好的移植性， 同时也可以⽤于测试验证其它⽅式获取的goid是否正确。
// 每个Go语⾔⽤户应该都知道panic函数。 调⽤panic函数将导致Goroutine异常，
// 如果panic在传递到Goroutine的根函数还没有被recover函数处理掉， 那么运⾏时将打印相关的异常和栈信息并退出Goroutine。

/*
panic("goid")

输出内容如下：
panic: goid

goroutine 1 [running]:
main.main()
        E:/gocode/crs/go_code/project03/main/chapter03/004getgoid/main.go:13 +0x40
这里可以看出  goroutine 1 其中数字 1 就是goid

但是如何才能在程序中获取panic的输出信息呢？ 其实上述信息只是当前函数调⽤栈帧的⽂字化描述， runtime.Stack函数提供了获取该信息的功能。

goroutine 1 [running]:
main.main()
        E:/gocode/crs/go_code/projec

因此从runtime.Stack获取的字符串中就可以很容易解析出goid信息。
 */
func main() {
	//panic("goid")

	var buf = make([]byte, 64)
	var stk = buf[:runtime.Stack(buf, false)]
	print(string(stk))

	fmt.Println(GetGoid())
}

func GetGoid() int64 {
	// [64]byte
	var buf = make([]byte, 64)
	var (
		//n = runtime.Stack(buf[:], false)
		n = runtime.Stack(buf, false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}