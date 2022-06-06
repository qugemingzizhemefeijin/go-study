package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	//n := rVal.Float() //error: call of reflect.Value.Float on int Value
	//fmt.Println("n=", n)
}

//通过反射，修改
//num int的值
//修改student的值
func reflectTest02(b interface{}) {
	//rType := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v, rVal Type %T, rVal Kind=%v\n", rVal, rVal, rVal.Kind()) //传递的指针，则Kind=ptr
	//需要将指针类型的转成值类型
	//rVal.SetInt(20) //error
	rVal.Elem().SetInt(20)
}

func main() {
	//1.reflect.Value.Kind，获取变量的类型，返回的是一个常量

	//2.Type是类型，Kind是类别，Type和Kind可能是相同的，也可能是不同的。
	// var num int = 10 num的Type是int,kind也是int
	// var stu Student stu的Type是(main.Student/modal.Student)，Kind是struct

	//3.通过反射可以让变量在interface{}和Relfect.Value之间相互交换。

	//4.使用反射的方式来获取变量的值(并返回对应的类型)，要求数据类型匹配，
	//  比如x是int，那么就应该使用reflect.Value(x).Int()，而不能会用其他的，否则报panic
	reflectTest01(1)

	//5.通过反射来修改变量，注意当使用SetXxx方法来设置需要通过对应的指针类型来完成，
	//  这样才能改变传入的变量的值，同时需要使用到reflect.Value.Elem()方法
	var num int = 10
	reflectTest02(&num)
	fmt.Println("num=", num)

	fmt.Println("exit success")
}
