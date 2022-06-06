package main

import (
	"fmt"
	"reflect"
)

//1.给你一个变量 var v float64 = 1.2，请使用反射来得到它的reflect.Value，然后获取对应的Type，Kind的值
//  并将reflect.Value转换成interface{}，再将interface{}转换成float64
func reflectTest01(v float64) {
	rVal := reflect.ValueOf(v)
	fmt.Printf("v type=%v, v kind=%v \n", rVal.Type(), rVal.Kind())

	vt := rVal.Interface()
	fmt.Printf("v interface=%v\n", vt)

	v1 := vt.(float64)
	fmt.Println("v value=", v1)
}

//以下代码报错了：error:reflect.Value.SetString using unaddressable value
func reflectTest02() {
	// var str string = "tom"
	// fs := reflect.ValueOf(str)
	// fs.SetString("jack") //这里会报错
	// fmt.Printf("%v\n", str)

	//改成这样子就可以了
	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack") //这里会报错
	fmt.Printf("%v\n", str)
}

func main() {
	var v float64 = 1.2
	reflectTest01(v)
	reflectTest02()
}
