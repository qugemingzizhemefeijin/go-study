package main

import "C"
import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var go_qsort_compare_info struct {
	base       unsafe.Pointer      // 排序数组的地址
	elemnum    int                 // 元素个数
	elemsize   int                 // 元素大小
	less       func(a, b int) bool // 比较函数
	sync.Mutex                     // 锁
}

//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	var (
		// array memory is locked
		base     = uintptr(go_qsort_compare_info.base)     // 数组的地址
		elemsize = uintptr(go_qsort_compare_info.elemsize) // 数组元素的大小
	)

	i := int((uintptr(a) - base) / elemsize)
	j := int((uintptr(b) - base) / elemsize)
	switch {
	case go_qsort_compare_info.less(i, j): // v[i] < v[j]
		return -1
	case go_qsort_compare_info.less(j, i): // v[i] > v[j]
		return +1
	default:
		return 0
	}
}

func Slice(slice interface{}, less func(a, b int) bool) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("qsort called with non-slice value of type %T", slice))
	}
	if sv.Len() == 0 {
		return
	}

	go_qsort_compare_info.Lock()
	defer go_qsort_compare_info.Unlock()

	defer func() {
		go_qsort_compare_info.base = nil
		go_qsort_compare_info.elemnum = 0
		go_qsort_compare_info.elemsize = 0
		go_qsort_compare_info.less = nil
	}()

	// baseMem = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	// baseMem maybe moved, so must saved after call C.fn
	go_qsort_compare_info.base = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	go_qsort_compare_info.elemnum = sv.Len()
	go_qsort_compare_info.elemsize = int(sv.Type().Elem().Size())
	go_qsort_compare_info.less = less
	C.qsort(
		go_qsort_compare_info.base,
		C.size_t(go_qsort_compare_info.elemnum),
		C.size_t(go_qsort_compare_info.elemsize),
		C.qsort_cmp_func_t(C._cgo_qsort_compare),
	)
}

// 改进： 消除⽤户对unsafe包的依赖
// func Slice(slice interface{}, less func(a, b int) bool)
func main() {
	// 为了避免在排序过程中， 排序数组的上下⽂信息 go_qsort_compare_info 被修改， 我们进⾏了全局加锁。
	// 因此⽬前版本的qsort.Slice函数是⽆法并发执⾏的。
	values := []int64{42, 9, 101, 95, 27, 25}
	Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values)
}
