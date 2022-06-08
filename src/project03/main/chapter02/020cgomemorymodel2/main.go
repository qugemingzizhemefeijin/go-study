package main

/*
goroutinue的栈因为空间不⾜的原因发⽣了扩展， 也就是导致了原来的Go语⾔内存被移动到了新的位置。
但是此时此刻C语⾔函数并不知道该Go语⾔内存已经移动了位置， 仍然⽤之前的地址来操作该内存——这将将导致内存越界。

借助C语⾔内存稳定的特性，在C语⾔空间先开辟同样⼤⼩的内存，然后将Go的内存填充到C的内存空间；返回的内存也是如此处理。

下⾯的例⼦是这种思路的具体实现：
*/

/*
#include <stdio.h>
#include <stdlib.h>

void printString(const char* s) {
	printf("%s", s);
}
 */
import "C"
import "unsafe"

// 在需要将Go的字符串传⼊C语⾔时，先通过 C.CString 将Go语⾔字符串对应的内存数据复制到新创建的C语⾔内存空间上。
// 例⼦的处理思路虽然是安全的，但是效率极其低下（因为要多次分配内存并逐个复制元素），同时也极其繁琐。
func printString(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	C.printString(cs)
}

// C临时访问传⼊的Go内存
func main() {
	s := "hello, world"
	printString(s)
}
