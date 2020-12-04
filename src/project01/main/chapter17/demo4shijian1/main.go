package main

import (
	"fmt"
	"reflect"
)

//Monster 这里申明了一个Monster结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

//Print  ...
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end----")
}

//GetSum ...
func (s Monster) GetSum(n1, n2 int) int {
	return 1 + n2
}

//Set ...
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

//TestStruct ...
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("epect struct")
		return
	}

	//获取字段数量
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		//typ.Field(i) 获取字段描述信息
		//val.Field(i) 获取字段值信息
		fmt.Printf("Field %d: 值为=%v structField=%v\n", i, val.Field(i), typ.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段有tag标签就显示，否则不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法(typ.NumMethod() == val.NumMethod())
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods, type methods %d \n", numOfMethod, typ.NumMethod())

	//Method的排序是按照方法名称的ASCII码来进行排序的，所以Print()方法，排列在第二个

	//var params []reflect.Value
	//这个是调用第二个方法
	val.Method(1).Call(nil)

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是 []reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果，返回的结果是 []reflect.Value
}

//使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
//func (v Value) Method(i int) Value	//默认按方法名排序对应i值，i从0开始
//func (v Value) Call(in []Value) []Value //传入参数，返回值是[]reflect.Value
func main() {
	//创建了一个Monster实例
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	//将Monster实例传递给TestStruct函数
	TestStruct(a)
	fmt.Println("exit success")
}
