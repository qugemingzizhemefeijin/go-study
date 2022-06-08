package main

// 在Go语⾔中，数组是⼀种值类型，⽽且数组的⻓度是数组类型的⼀个部分。Go语⾔字符串对应⼀段⻓度确定的只读byte类型的内存。Go语⾔的切⽚则是⼀个简化版的动态数组。
// Go语⾔和C语⾔的数组、 字符串和切⽚之间的相互转换可以简化为Go语⾔的切⽚和C语⾔中指向⼀定⻓度内存的指针之间的转换。

// C.CString 针对输⼊的Go字符串， 克隆⼀个C语⾔格式的字符串，返回的字符串由C语⾔的 malloc 函数分配， 不使⽤时需要通过C语⾔的 free 函数释放。
// C.CBytes 函数的功能和 C.CString 类似， ⽤于从输⼊的Go语⾔字节切⽚克隆⼀个C语⾔版本的字节数组， 同样返回的数组需要在合适的时候释放。
// C.GoString ⽤于将从NULL结尾的C语⾔字符串克隆⼀个Go语⾔字符串。
// C.GoStringN 是另⼀个字符数组克隆函数。
// C.GoBytes ⽤于从C语⾔数组， 克隆⼀个Go语⾔字节切⽚。

// 该组辅助函数都是以克隆的⽅式运⾏。 当Go语⾔字符串和切⽚向C语⾔转换时， 克隆的内存由C语⾔的 malloc 函数分配， 最终可以通过 free 函数释放。
// 当C语⾔字符串或数组向Go语⾔转换时，克隆的内存由Go语⾔分配管理。

// 当C语⾔字符串或数组向Go语⾔转换时， 克隆的内存由Go语⾔分配管理。
// 克隆⽅式实现转换的优点是接⼝和内存管理都很简单， 缺点是克隆需要分配新的内存和复制操作都会导致额外的开销。

/*
// 在 reflect 包中有字符串和切⽚的定义：
type StringHeader struct {
	Data uintptr
	Len int
}

type SliceHeader struct {
	Data uintptr
	Len int
	Cap int
}
 */

/*
static char arr[10];
static char *s = "Hello";
 */
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 如果不希望单独分配内存， 可以在Go语⾔中直接访问C语⾔的内存空间：
	// 通过 reflect.SliceHeader 转换
	var arr0 []byte	// 定义一个nil切片
	var arr0Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&arr0))
	arr0Hdr.Data = uintptr(unsafe.Pointer(&C.arr[0])) //将切片的Data属性指针指向C的地址
	arr0Hdr.Len = 10 // 长度改为10
	arr0Hdr.Cap = 10 // cap改为10

	// 通过切⽚语法转换
	arr1 := (*[31]byte)(unsafe.Pointer(&C.arr[0]))[:10:10]

	var s0 string
	var s0Hdr = (*reflect.StringHeader)(unsafe.Pointer(&s0))
	s0Hdr.Data = uintptr(unsafe.Pointer(C.s))
	s0Hdr.Len = int(C.strlen(C.s)) // 不知道为什么这里编译不通 could not determine kind of name for C.strlen。到时候看一下。。

	sLen := int(C.strlen(C.s))
	s1 := string((*[31]byte)(unsafe.Pointer(&C.s[0]))[:sLen:sLen])

	fmt.Println(arr1)
	fmt.Println(s1)
}
