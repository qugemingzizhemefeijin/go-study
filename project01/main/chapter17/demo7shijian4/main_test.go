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

func TestReflectStruct(t *testing.T) {
	var (
		model *User
		sv    reflect.Value
	)
	model = &User{}
	sv = reflect.ValueOf(model) //获取类型*User
	//这里是个指针类型了
	t.Log("reflect.ValueOf", sv.Kind().String()) //ptr

	//这里将指针指向的内存地址给sv变量，就可以实际操作对象的内存空间了
	sv = sv.Elem()
	t.Log("reflect.ValueOf.Elem", sv.Kind().String()) //struct

	sv.FieldByName("UserID").SetString("123456")
	sv.FieldByName("Name").SetString("小明")

	t.Log("model", model)
}

func main() {
	fmt.Println("exit success")
}
