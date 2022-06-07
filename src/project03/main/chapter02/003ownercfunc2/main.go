package main

//void SayHello(const char* s);
import "C"

// 在CGO部分先声明 SayHello 函数
func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
