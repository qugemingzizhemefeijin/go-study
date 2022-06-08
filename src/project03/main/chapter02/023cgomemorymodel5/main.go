package main

// 导出 C 函数不能返回 Go 内存
// 在 Go 语言中，Go 是从一个固定的虚拟地址空间分配内存。而 C 语言分配的内存则不能使用 Go 语言保留的虚拟内存空间。
// 在 CGO 环境，Go 语言运行时默认会检查导出返回的内存是否是由 Go 语言分配的，如果是则会抛出运行时异常。

/*
extern int* getGoPtr();

static void Main() {
	int* p = getGoPtr();
	*p = 42;
}
 */
import "C"

//export getGoPtr
func getGoPtr() *C.int {
	return new(C.int)
}

// cgo 默认对返回结果的指针的检查是有代价的，特别是 cgo 函数返回的结果是一个复杂的数据结构时将花费更多的时间。
// 如果已经确保了 cgo 函数返回的结果是安全的话，可以通过设置环境变量 GODEBUG=cgocheck=0 来关闭指针检查行为。

// 关闭 cgocheck 功能后再运行下面的代码就不会出现下面的异常的。但是要注意的是，
// 如果 C 语言使用期间对应的内存被 Go 运行时释放了，将会导致更严重的崩溃问题。
// cgocheck 默认的值是 1，对应一个简化版本的检测，如果需要完整的检测功能可以将 cgocheck 设置为 2。
func main() {
	// 下面是 CGO 运行时异常的例子
	// 其中 getGoPtr 返回的虽然是 C 语言类型的指针，但是内存本身是从 Go 语言的 new 函数分配，也就是由 Go 语言运行时统一管理的内存。
	// 然后我们在 C 语言的 Main 函数中调用了 getGoPtr 函数，此时默认将发送运行时异常
	C.Main()

	// 异常说明 cgo 函数返回的结果中含有 Go 语言分配的指针。指针的检查操作发生在 C 语言版的 getGoPtr 函数中，
	// 它是由 cgo 生成的桥接 C 语言和 Go 语言的函数。
	/*
	panic: runtime error: cgo result has Go pointer

	goroutine 1 [running]:
	main._cgoexpwrap_bc6b838f5e6e_getGoPtr.func1(0xc00003bdb8)
	        _cgo_gotypes.go:60 +0x41
	main._cgoexpwrap_bc6b838f5e6e_getGoPtr(0xc000094010)
	        _cgo_gotypes.go:62 +0x86
	main._Cfunc_Main()
	        _cgo_gotypes.go:43 +0x48
	main.main()
	*/
}
