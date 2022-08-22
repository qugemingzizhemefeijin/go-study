package main

import "fmt"

// 上一节使用到了typefunc的知识点，这一节加深一下印象

// 在go语言中，type可以定义任何自定义的类型
// 所以func也是可以作为类型自定义的，type myFunc func(int) int，意思是自定义了一个叫myFunc的函数类型，这个函数的签名必须符合输入为int，输出为int
// 已知，相同底层类型的变量之间是可以相互转换的，例如从一个取值范围小的int16转为取值范围大的int32
// 所以，自定义的 myInt 和 int 之间也是可以转换的

type myInt int

// 有什么场景需要自己定义一个myInt出来？就是重载原来的int类型，并自定义新的方法
func (mi myInt) IsZero() bool {
	return mi == 0
}

// 同理，myFunc 也是可以将签名为 func(int) int 的函数转换成 myFunc 类型
type myFunc func(int) int

// 定义函数类型也可以自定义方法
func (mf myFunc) sumx(a,b int) int {
	c := a + b
	return mf(c)
}

func sumx10(num int) int {
	fmt.Println(num*10)
	return num * 10
}

// 一个自定义函数类型的变量拥有了一个sum函数，有什么实际用途？
func (f myFunc) sum(a, b int) int {
	res := a + b
	return f(res)
}

func sum10(num int) int {
	return num * 10
}

func sum100(num int) int {
	return num * 100
}

// 假如handlerSum是一种特殊的sum算法，但是又有一部分的计算是可以通过外部自定义函数来干预的，那么使用这种方式就很合适。
func handlerSum(handler myFunc, a, b int) int {
	res := handler.sum(a, b)
	fmt.Println(res)
	return res
}

// 再进一步，如何使得handlerSum函数更抽象化？我必须传递一个myFunc类型的变量参数进来吗？
// 参数是一个interface呢，一个拥有sum方法的interface是不是更通用？
type sumable interface {
	sum(int, int) int
}

// icansum结构体继承sumable接口
type icansum struct {
	name string
	res int
}

func (ics *icansum) sum(a, b int) int {
	ics.res = a + b
	return ics.res
}

// handler只要是继承了sumable接口的任何变量都行，我只需要你提供sum函数就好
func handlerSum2(handler sumable, a, b int) int {
	res := handler.sum(a, b)
	fmt.Println(res)
	return res
}

func main() {
	var a int
	a = 2
	b := myInt(a)

	fmt.Println(b) // 2

	newFunc := myFunc(sumx10) // 实际就是将sum10改成myFunc类型，然后通过别名来调用sum10
	newFunc(10) // 100

	b = 0
	fmt.Println(b.IsZero())        // true

	newFunc1 := myFunc(sum10)
	newFunc2 := myFunc(sum100)

	handlerSum2(newFunc1, 1, 1)    // 20
	handlerSum2(newFunc2, 1, 1)    // 200

	ics := &icansum{"I can sum", 0}
	handlerSum2(ics, 1, 1)         // 2
}