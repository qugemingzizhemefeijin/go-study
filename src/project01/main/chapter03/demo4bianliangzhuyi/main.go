package main
import "fmt"

//变量使用的注意事项
func main(){
	var i int = 10
	i = 30
	i = 50

	fmt.Println("i=", i)

	//i = 1.2 //int不能赋值给float

	// +号的使用
	var c , d = 10, 20
	var e = c + d
	fmt.Println("c = ", e)

	var n , m = "Hello", "World"
	var z = n + m
	fmt.Println("name=", z)
}