package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest01(b interface{}) {
	//通过反射获取传入的变量的type、kind、值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType = ", rTyp)

	//2.获取到reflect.Value类型
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)

	fmt.Printf("rVal = %v, rVal type=%T\n", rVal, rVal)

	//3.将rVal类型转成interface{}
	iv := rVal.Interface()
	fmt.Printf("iv = %v, iv type = %T\n", iv, iv)

	//4.家那个interface{}通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Printf("num2 = %v, num2 type = %T\n", num2, num2)
}

//Student ...
type Student struct {
	Name string
	Age  int
}

//抓们演示对结构体的反射
func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType = ", rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v, rVal type=%T\n", rVal, rVal)

	//获取 变量对应的Kind
	//(1) rVal.Kind() ===>
	//(2) rTyp.Kind() ===>
	fmt.Printf("kind1 = %v, kind2 = %v\n", rVal.Kind(), rTyp.Kind())

	iv := rVal.Interface()
	fmt.Printf("iv = %v, iv type = %T\n", iv, iv)

	//将interface{}通过断言转成需要的类型
	// 这里也可以用是switch的断言形式来做的更加灵活
	// stu, ok := iv.(Student)
	// if !ok {
	// 	fmt.Println("duanyan error")
	// } else {
	// 	fmt.Println("stu=", stu)
	// }
	switch iv.(type) {
	case Student:
		fmt.Println("stu = ", iv.(Student))
	default:
		fmt.Println("not found type")
	}
}

func main() {
	//1.请编写一个案例，演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作。
	//先定义一个int
	num := 100
	reflectTest01(num)

	//2.请编写一个案例，演示对(结构体类型、interface{}、reflect.Value)进行反射的基本操作
	stu := Student{
		Name: "Tom",
		Age:  20,
	}
	reflectTest02(stu)
}
