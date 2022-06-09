package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type X struct {
	a int
}

type Y struct {
	b uint32
}

// Go语⾔中， 数组或数组对应的切⽚都不再是指针类型， 因此我们也就⽆法直接实现不同类型的切⽚之间的转换
// 不过Go语⾔的reflect包提供了切⽚类型的底层结构， 再结合前⾯讨论到不同类型之间的指针转换技术就可以实现 []X 和 []Y 类型的切⽚转换：
func main() {
	var p []X
	var q []Y

	fmt.Println(p, q)

	pHdr := (*reflect.SliceHeader)(unsafe.Pointer(&p))
	qHdr := (*reflect.SliceHeader)(unsafe.Pointer(&q))

	pHdr.Data = qHdr.Data
	// 不知道为什么例子是错误的。。。。 invalid operation: qHdr.Len * unsafe.Sizeof(q[0]) (mismatched types int and uintptr)
	//pHdr.Len = qHdr.Len * unsafe.Sizeof(q[0]) / unsafe.Sizeof(p[0])
	//pHdr.Cap = qHdr.Cap * unsafe.Sizeof(q[0]) / unsafe.Sizeof(p[0])

	fmt.Println(p, q)
}
