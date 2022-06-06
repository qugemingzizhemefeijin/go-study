package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

// 为了安全，一般切片不可以直接转换类型。但是有时候可以将float+4转换为int排序，提高排序性能，这个时候就可以使用骚操作来强制转换

var a1 = []float64{4, 2, 5, 7, 2, 1, 88, 1}
var a2 = []float64{4, 2, 5, 7, 2, 1, 88, 1}

// SortFloat64FastV1 排序1
func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序
	sort.Ints(b)

	for _, v := range a {
		fmt.Println(v)
	}
}

// SortFloat64FastV2 排序2
func SortFloat64FastV2(a []float64) {
	// 通过reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(c)

	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	// 先将切片数据的开始地址转换为一个较大的数组的指针，然后对数组指针对应的数组重新做切片操作。
	// 中间需要unsafe.Pointer来连接两个不同类型的指针传递
	SortFloat64FastV1(a1)
	fmt.Println("=================")
	// 分别取到两个不同类型的切片头信息指针，任何类型的切片头部信息底层都是对应reflect.SliceHeader结构。
	// 然后通过更新结构体方式来更新切片信息，从而实现a对应的[]float64切片到c对应的[]int类型切片的转换。
	SortFloat64FastV2(a2)
}
