package main

import "fmt"

// main				pkg1			    pkg2
// import pkg1		import pkg2 	    const
// const  <----|	const      <---|    var
// var		   |	var            |--- init()
// init()	   |	init()
// main()      ----	init()
//

// Add 具名函数
func Add(a, b int) int {
	return a + b
}

// Add2 匿名函数
var Add2 = func(a, b int) int {
	return a + b
}

// Swap 多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// Sum 可变数量的参数，more对应 []int切片类型，可变参数必须是最后出现
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

// Print 当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不用的结果
func Print(a ...interface{}) {
	fmt.Println(a...)
}

// Find 不仅函数的参数可以有名字，也可以给函数的返回值命名
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

// Inc refer语句延迟执行了一个匿名函数，这个匿名函数捕获了外部函数的局部变量v，所以也叫做闭包
// 闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问。
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

// 如果以切⽚为参数调⽤函数时， 有时候会给⼈⼀种参数采⽤了传引⽤的⽅式的假象： 因为在被调⽤函数内部可以修改传⼊的切⽚的元素。
// 其实， 任何可以通过函数参数修改调⽤参数的情形，都是因为函数参数中显式或隐式传⼊了指针参数。
// 函数参数传值的规范更准确说是只针对数据结构中固定的部分传值，
// 例如字符串或切⽚对应结构体中的指针和字符串⻓度结构体传值， 但是并不包含指针间接指向的内容。

// Go语⾔函数的递归调⽤深度逻辑上没有限制， 函数调⽤的栈是不会出现溢出错误的， 因为Go语⾔运⾏时会根据需要动态地调整函数栈的⼤⼩。
// 每个goroutine刚启动时只会分配很⼩的栈（4或8KB， 具体依赖实现） ， 根据需要动态调整栈的⼤⼩， 栈最⼤可以达到GB级
func main() {
	a := []interface{}{123, "abc"}

	Print(a...) // 123 abc
	Print(a)    // [123 abc] 等价于直接调用 Print([]interface{}{123, "abc"})

	var b = Inc() // 看函数是返回42，实际b = 43
	Print(b)

	Print("===================")

	// 闭包这种引用方式访问外部变量的行为可能会导致一些隐含的问题
	for i := 0; i < 3; i++ {
		defer func() { Print(i) }() // 因为defer语句延迟执行的函数引用的都是同一i迭代变量，在循环结束后这个变量的值为3，因此输出的均为3
	}

	// 解决办法如下：
	// 第一种方式是在循环体内定义一个局部变量，这样每次迭代defer语句的闭包函数捕获的都是不同的变量。
	// 第二种是将迭代变量通过闭包函数的参数传入，defer语句会马上对调用参数求值。

	for i := 0; i < 3; i++ {
		j := i // 定义一个循环体内局部变量i
		defer func() { Print(j) }()
	}

	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) { Print(i) }(i)
	}
}
