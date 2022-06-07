package main

//static const char* cs = "hello"
import "C"

import "./cgo_helper"

func main() {
	// 这段代码是不能正常⼯作的， 因为当前main包引⼊的 C.cs 变量的类型是当前 main 包的cgo构造的虚拟的C包下的 *char 类型
	//（具体点是 *C.char ， 更具体点是 *main.C.char ） ， 它和cgo_helper包引⼊的 *C.char 类型（具体点是 *cgo_helper.C.char ） 是不同的。
	// 在Go语⾔中⽅法是依附于类型存在的， 不同Go包中引⼊的虚拟的C包的类型却是不同的（ main.C 不等 cgo_helper.C ） ，
	// 这导致从它们延伸出来的Go类型也是不同的类型（ *main.C.char 不等 *cgo_helper.C.char ） ， 这最终导致了前⾯代码不能正常⼯作。

	// 有Go语⾔使⽤经验的⽤户可能会建议参数转型后再传⼊。但是这个⽅法似乎也是不可⾏的，因为cgo_helper.PrintCString的参数是它⾃身包引⼊的*C.char类型，
	// 在外部是⽆法直接获取这个类型的。 换⾔之， ⼀个包如果在公开的接⼝中直接使⽤了 *C.char 等类似的虚拟C包的类型， 其它的Go包是⽆法直接使⽤这些类型的，
	// 除⾮这个Go包同时也提供了 *C.char 类型的构造函数。 因为这些诸多因素， 如果想在go test环境直接测试这些cgo导出的类型也会有相同的限制。
	cgo_helper.PrintCString(C.cs)
}
