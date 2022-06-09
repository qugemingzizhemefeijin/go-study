package main

/*
qsort 快速排序函数有 <stdlib.h> 标准库提供，函数的声明如下：

void qsort(void* base, size_t num, size_t size, int (*cmp)(const void*, const void*));

base 参数是要排序数组的首个元素的地址
num 是数组中元素的个数
size 是数组中每个元素的大小
cmp 比较函数，用于对数组中任意两个元素进行排序

C语言的例子：
#include <stdio.h>
#include <stdlib.h>

#define DIM(x) (sizeof(x)/sizeof((x)[0]))

static int cmp(const void* a, const void* b) {
	const int* pa = (int*)a;
	const int* pb = (int*)b;
	return *pa - *pb;
}

int main() {
	int values[] = { 42, 8, 109, 97, 23, 25};
	int i;

	qsort(values, DIM(values), sizeof(values[0]), cmp);

	for(i = 0; i < DIM(values); i++) {
		printf ("%d",values[i]);
	}
	return 0;
}
*/

/*
#include <stdlib.h>


// CGO 语言不好直接表达 C 语言的函数类型，因此在 C 语言空间将比较函数类型重新定义为一个 qsort_cmp_func_t 类型。
typedef int (*qsort_cmp_func_t)(const void* a, const void* b);

extern int go_qsort_compare(void* a, void* b);
 */
import "C"
import (
	"fmt"
	"unsafe"
)

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

type CompareFunc C.qsort_cmp_func_t

// 在CGO的内部机制⼀节中我们已经提过， 虚拟的C包下的任何名称其实都会被映射为包内的私有名字。
// ⽐如 C.size_t 会被展开为 _Ctype_size_t ， C.qsort_cmp_func_t 类型会被展开为 _Ctype_qsort_cmp_func_t 。
// 被CGO处理后的Sort函数的类型如下：
/*
func Sort(base unsafe.Pointer, num, size _Ctype_size_t,cmp _Ctype_qsort_cmp_func_t,)

这样将会导致包外部⽤于⽆法构造 _Ctype_size_t 和 _Ctype_qsort_cmp_func_t 类型的参数⽽⽆法使⽤Sort函数。
因此， 导出的Sort函数的参数和返回值要避免对虚拟C包的依赖。

不能直接使用如下：
func Sort(base unsafe.Pointer, num, size C.size_t,cmp C.qsort_cmp_func_t,) {
	C.qsort(base, num, size, cmp)
}

需要转成下面依赖go类型的数据
 */
func Sort(base unsafe.Pointer, num, size int, cmp CompareFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(cmp))
}

// ⽤Go语⾔将qsort函数重新包装为 qsort.Sort 函数
func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	Sort(unsafe.Pointer(&values[0]),
		len(values), int(unsafe.Sizeof(values[0])),
		CompareFunc(C.go_qsort_compare),
	)
	fmt.Println(values)
}