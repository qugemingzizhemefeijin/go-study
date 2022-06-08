package main

/*
C语⾔空间的内存是稳定的， 只要不是被⼈为提前释放， 那么在Go语⾔空间可以放⼼⼤胆地使⽤。 在Go语⾔访问C语⾔内存是最简单的情形， 我们在之前的例⼦中已经⻅过多次。

因为Go语⾔实现的限制， 我们⽆法在Go语⾔中创建⼤于2GB内存的切⽚（具体请参考makeslice实现代码）。
不过借助cgo技术， 我们可以在C语⾔环境创建⼤于2GB的内存， 然后转为Go语⾔的切⽚使⽤。
 */

/*
#include <stdlib.h>

void* makeslice(size_t memsize) {
	return malloc(memsize);
}
 */
import "C"
import "unsafe"

func makeByteSlize(n int) []byte {
	p := C.makeslice(C.size_t(n))
	return ((*[1 << 31]byte)(p))[0:n:n]
}

func freeByteSlice(p []byte) {
	C.free(unsafe.Pointer(&p[0]))
}

// 例⼦中我们通过makeByteSlize来创建⼤于4G内存⼤⼩的切⽚， 从⽽绕过了Go语⾔实现的限制（需要代码验证） 。
// ⽽freeByteSlice辅助函数则⽤于释放从C语⾔函数创建的切⽚。
// 因为C语⾔内存空间是稳定的， 基于C语⾔内存构造的切⽚也是绝对稳定的， 不会因为Go语⾔栈的变化⽽被移动。
func main() {
	s := makeByteSlize(1<<32+1)
	s[len(s)-1] = 255
	print(s[len(s)-1])
	freeByteSlice(s)
}
