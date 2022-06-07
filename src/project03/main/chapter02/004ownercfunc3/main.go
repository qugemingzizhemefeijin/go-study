package main

//void SayHello(char* s);
import "C"

import "fmt"

//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
