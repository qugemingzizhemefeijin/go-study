package main

/*
#include <stdio.h>
void printint(int v) {
	printf("printint: %d\n", v);
}
*/
import "C"

// 如果在Go代码中出现了 import "C" 语句则表示使⽤了CGO特性， 紧跟在这⾏语句前⾯的注释是⼀种特殊语法，
// ⾥⾯包含的是正常的C语⾔代码。 当确保CGO启⽤的情况下， 还可以在当前⽬录中包含C/C++对应的源⽂件。
func main() {
	v := 42
	C.printint(C.int(v))
}
