package main

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
extern int _cgo_qsort_compare(void* a, void* b);
 */
import "C"
import (
	"fmt"
	"sort"
	"sync"
	"unsafe"
)

// 缺点是需要一个全局的变量来维持锁以及fn比较函数
var go_qsort_compare_info struct {
	fn func(a, b unsafe.Pointer) int
	sync.Mutex
}

// 其中导出的C语⾔函数 _cgo_qsort_compare 是公⽤的qsort⽐较函数， 内部通过 go_qsort_compare_info.fn 来调⽤当前的闭包⽐较函数。
//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	return C.int(go_qsort_compare_info.fn(a, b))
}

func Sort(base unsafe.Pointer, num, size int, cmp func(a, b unsafe.Pointer) int) {
	go_qsort_compare_info.Lock()
	defer go_qsort_compare_info.Unlock()

	go_qsort_compare_info.fn = cmp

	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(C._cgo_qsort_compare))
}

// 改进： 闭包函数作为⽐较函数
// 我们也尝试将C语⾔的qsort函数包装为以下格式的Go语⾔函数：
// func Sort(base unsafe.Pointer, num, size int, cmp func(a, b unsafe.Pointer) int)
// 闭包函数⽆法导出为C语⾔函数， 因此⽆法直接将闭包函数传⼊C语⾔的qsort函数。
// 为此我们可以⽤Go构造⼀个可以导出为C语⾔的代理函数， 然后通过⼀个全局变量临时保存当前的闭包⽐较函数。
func main() {
	// 这个是go内置的排序
	values := []int32{42, 9, 101, 95, 27, 25}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values)

	values2 := []int32{42, 9, 101, 95, 27, 25}
	// 现在排序不再需要通过CGO实现C语⾔版本的⽐较函数了， 可以传⼊Go语⾔闭包函数作为⽐较函数。
	// 但是导⼊的排序函数依然依赖unsafe包， 这是违背Go语⾔编程习惯的。
	Sort(unsafe.Pointer(&values2[0]), len(values2), int(unsafe.Sizeof(values2[0])),
		func(a, b unsafe.Pointer) int {
			pa, pb := (*int32)(a), (*int32)(b)
			return int(*pa - *pb)
		},
	)
	fmt.Println(values2)
}