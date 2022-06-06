package main

import (
	"fmt"
)

func test03(n1 *int) {
	fmt.Printf("n1的地址 %v\n", &n1)
	*n1 = *n1 + 10
	fmt.Println("test03() n1=", *n1)
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funVar func(int, int) int, num1 int, num2 int) int {
	return funVar(num1, num2)
}

type myFunType func(int, int) int

func myFun2(funVar myFunType, num1 int, num2 int) int {
	return funVar(num1, num2)
}

//支持对返回值命名，就是省点事情
func cal(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func sum(n1 int, args ...int) int {
	var s int = n1
	//遍历args
	for i := 0; i < len(args); i++ {
		s += args[i]
	}
	return s
}

//这种n1后面不定义类型是可以的，n1和n2就是一样的类型了，跟定义变量 var n1,n2 float = 1.0, 2.0效果是一样的
func test(n1, n2 float32) float32 {
	return n1 + n2
}

func main() {
	//1.函数的形参列表可以是多个，返回值列表也可以是多个
	//2.形参列表和返回值列表的数据类型可以是值类型和引用类型
	//3.函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写该函数可以被本包文件和其他包文件使用，类似public，
	//  首字母小写，只能被本包文件使用，其他包文件不能使用，类似private
	//4.函数中的变量是局部的
	//5.基本数据类型和数组默认都是值传递，即进行值拷贝。在函数内修改，不会影响到原来的值
	//6.如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量。
	num := 20
	fmt.Printf("num address %v\n", &num)
	test03(&num)
	fmt.Println("main() num=", num)
	//7.Go函数不支持重载。
	//8.在Go只能跟，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量。通过该变量可以对函数调用。
	a := getSum
	fmt.Printf("a的类型%T，getSum类型是%T\n", a, getSum)
	res := a(10, 40) //等价 res := getSum(10, 40)
	fmt.Println("res=", res)
	//9.函数既然是一种数据类型，因为在Go中，函数可以作为形参，并且调用。
	fmt.Println("myFun=", myFun(a, 10, 40))
	//10.为了简化数据类型定义，Go支持自定义数据类型
	// type myInt int  //这是myInt就等价int来使用
	// type mySum func(int,int) int //这时mySum就等价一个函数类型
	type myInt int //给int取了别名，在go中myInt和int虽然都是int类型，但是go认为还是两个不同的类型
	var num1 myInt
	var num2 int
	num1 = 40
	//num2 = num1 //虽然都是int，但是这里会报错的
	num2 = int(num1) //这里需要进行显示转换一下
	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	//再举一个案例
	//type myFunType func(int, int) int //这时 myFun 就是函数类型
	fmt.Println("myFun=", myFun2(getSum, 10, 40))
	//11.支持对函数返回值命名
	x1, x2 := cal(30, 20)
	fmt.Printf("x1=%d,x2=%d\n", x1, x2)
	//12.实用 _标识符，忽略返回值
	x3, _ := cal(20, 30)
	fmt.Println("x3=", x3)
	//13.Go支持可变参数，支持0到多个参数 args... int，支持1到多个n1 int, args... int，可变参数需要放在形参列表的最后
	//   args是slice切片，通过args[index] 可以反问到各个值
	fmt.Println("sum=", sum(10, 20, 30, 40, 50))
}
