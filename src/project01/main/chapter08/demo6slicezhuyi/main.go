package main

import (
	"fmt"
)

//注意点
func main() {
	//1.切片初始化时，var slice = arr[startIndex:endIndex]，说明：从arr数组下标为startIndex，取到下标为endIndex的元素
	//2.切片初始化时，仍然不能越界。范围在[0-len(arr)]之间，但是可以动态增长
	//  1)var slice = arr[0:end] 可以简写var slice = arr[:end]
	//  2)var slice = arr[start:len(arr)] 可以简写 var slice = arr[start:]
	//  3)var slice = arr[0:len(arr)]可以简写 var slice = arr[:]
	//3.cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
	//4.切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组或者make一个空间供切片来使用
	//5.切片可以继续切片
	//6.用append内置函数，可以切切片进行动态追加【经测试，追加之后，切片指向的内存地址和原数组的内存地址就不一样了】
	//  注意：数组变量的指针指向数组的第一个元素的地址，而slice的指针与其第一个元素的地址不是一个地址；
	//  append的扩容原理
	//  1)切片append操作的本事就是对数组扩容
	//  2)go底层会创建一下新的数组newArr(安装扩容后大小)
	//  3)将slice原来包含的元素拷贝到新的数组newArr
	//  4)slice重新引用到newArr
	//  5)注意newArr是在底层来维护的，程序员不可见
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	var slice = arr[:]
	fmt.Printf("arr=%v,arr address=%p, arr[0] address=%p\n", arr, &arr, &arr[0])
	fmt.Printf("追加之前:slice=%v,len=%v,cap=%v,slice address=%p\n", slice, len(slice), cap(slice), &slice[0])
	//追加具体的元素
	slice = append(slice, 10, 20, 30)
	//fmt.Println("arr=", arr)
	fmt.Printf("追加之后:slice=%v,len=%v,cap=%v,slice address=%p\n", slice, len(slice), cap(slice), &slice)

	//在切片上追加切片
	var a = []int{100, 200}
	slice = append(slice, a...)
	//fmt.Println("arr=", arr)
	fmt.Printf("追加之后:slice=%v,len=%v,cap=%v,slice address=%p\n", slice, len(slice), cap(slice), &slice[0])

	slice[0] = 999
	//fmt.Println("arr=", arr)
	fmt.Printf("修改之后:slice=%v,len=%v,cap=%v,slice address=%p\n", slice, len(slice), cap(slice), &slice[1])

	//7.切片的拷贝操作，使用copy内置函数完成拷贝
	slice = append(slice, 888)
	fmt.Println("orignSlice=", slice)

	var cloneSlice = make([]int, 10)
	fmt.Println("cloneSlice=", cloneSlice)
	copy(cloneSlice, slice)
	fmt.Println("cloneSlice=", cloneSlice)
}
