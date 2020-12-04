package main

import (
	"fmt"
	"reflect"
	"testing"
)

//User ...
type User struct {
	UserID string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *User
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)                  //获取类型*user
	t.Log("reflect.TypeOf", st.Kind().String()) //ptr

	st = st.Elem()                                   //这个是获取User结构体信息
	t.Log("reflect.TypeOf.Elem", st.Kind().String()) //struct

	elem = reflect.New(st)                                 // New返回一个Value类型值，该值持有一个指向类型为typ的新申请的零值的指针
	t.Log("reflect.New", elem.Kind().String())             //ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) //struct

	//model就是创建的user结构体变量(实例)
	model = elem.Interface().(*User) //model是*User，它的指向和elem是一样的，其实就是都指向的同一个指针变量
	elem = elem.Elem()               //取得elem指向的值，其实就是获取New出来真正的实例对象的空间地址

	elem.FieldByName("UserID").SetString("123456") //赋值...
	elem.FieldByName("Name").SetString("狐狸精")

	t.Log("model model.Name", model, model.Name)
}

func TestReflectStruct(t *testing.T) {
	var (
		model User
		st    reflect.Type
		elem  reflect.Value
	)

	st = reflect.TypeOf(model)                  //获取类型user
	t.Log("reflect.TypeOf", st.Kind().String()) //strct

	elem = reflect.New(st)                                 // New返回一个Value类型值，该值持有一个指向类型为typ的新申请的零值的指针
	t.Log("reflect.New", elem.Kind().String())             //ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) //struct

	//model就是创建的user结构体变量(实例)
	elem = elem.Elem()              //取得elem指向的值，其实就是获取New出来真正的实例对象的空间地址
	model = elem.Interface().(User) //这里如果赋值的话，model输出的是空，因为结构体是复制传递

	elem.FieldByName("UserID").SetString("888888") //赋值...
	elem.FieldByName("Name").SetString("黄鼠狼精")

	model = elem.Interface().(User) //在这里赋值的话，就没问题了

	t.Log("model model.Name", model, model.Name)
}

//使用反射创建并操作结构体
func main() {

	fmt.Println("exit success")
}
