package main

import (
	"fmt"
)

func test(b byte) byte {
	return b + 1
}

//编写一个程序，接收一个字符，a,b,c,d,e,f,g，a表示星期一，b表示星期二...根据用户呃输入显示相应的信息。要求使用switch语句完成
func main() {
	// var key byte
	// fmt.Println("ing输入一个字符 a b c d e f g")
	// fmt.Scanf("%c", &key)

	// switch test(key) + 1 {
	// case 'a':
	// 	fmt.Println("星期一，猴子穿新衣")
	// case 'b':
	// 	fmt.Println("星期二，猴子当小二")
	// case 'c':
	// 	fmt.Println("星期三，猴子爬雪山")
	// case 'd':
	// 	fmt.Println("星期四，猴子风雪寺")
	// case 'e':
	// case 'f':
	// case 'g':
	// default:
	// 	fmt.Println("输入有误...")
	// }

	//switch后也可以不带表达式，类似if -- else分支来使用了
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配到")
	}

	//case中也可以对age范围进行判断
	var score int = 10
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score >= 70 && score <= 90:
		fmt.Println("优良")
	default:
		fmt.Println("及格")
	}

	//switch后也可以直接定义/申明一个变量，分号结束，不推荐
	switch grade := 90; {
	case grade > 90:
		fmt.Println("90")
	default:
		fmt.Println("牛逼")
	}

	//switch穿透fallthrough，如果在case语句块后增加fallthrough，则会继续执行下一个case，也就switch穿透
	var i = 10
	switch i {
	case 10:
		fmt.Println("OK1")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("OK2")
	case 30:
		fmt.Println("OK3")
	default:
		fmt.Println("没有匹配")
	}

	//Type Switch: switch语句还可以被用于type-switch来判断某个interface变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型~：%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知类型")
	}
}
