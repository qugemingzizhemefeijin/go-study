package main

import (
	"fmt"
	"project01/main/chapter03/demo15globalvar/model"
)

func main() {
	//使用utils.go的heroName变量， 包名.标识符(函数、变量等)
	fmt.Println(model.HeroName)
}
