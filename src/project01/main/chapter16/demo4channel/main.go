package main

import (
	"fmt"
)

//Cat ...
type Cat struct {
	Name string
	Age  int
}

func main() {
	//1.channel本质就是一个数据结构-队列
	//2.数据是先进先出
	//3.线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
	//4.channel是有类型的，一个string的channel只能存放string类型数据

	//定义/声明channel
	//var 变量名 chan 数据类型
	//举例：
	//var intChan chan int (intChan用于存放int数据)
	//var mapChan chan map[int]string (mapChan用于存放map[int]string类型)
	//var perChan chan Person
	//var perChan2 chan *Person

	//说明：
	//1.channel是引用类型
	//2.channel必须初始化才能写入数据，即make后才能使用
	//3.管道是有类型的，intChan只能写入整个int

	//注意事项：
	//1.channel只能存放指定类型
	//2.channel的数据放满后，就不能再放入了
	//3.如果从channel中取出数据后，可以继续放入
	//4.在没有使用协程的情况下，如果channel数据取完了，再取，就会报deadlock

	//演示
	//1. 创建一个可以存放3个int类型的管道
	var intChan chan int = make(chan int, 3)

	//2. 看看intChan是什么？
	fmt.Printf("intChan 的值=%v, intChan本身的地址=%v\n", intChan, &intChan)

	//3. 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	//<- intChan 获取数据，丢弃掉
	intChan <- 20
	//intChan <- 30	//注意点，当我们给管道写入数据时，不能超过其容量，否则会报DeadLock，因为没有其他地方来读取了

	//4. 看看管道呃长度和cap(容量)
	fmt.Printf("channel len=%v, cap=%v \n", len(intChan), cap(intChan))

	//5. 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v, cap=%v \n", len(intChan), cap(intChan))

	//6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告deadlock错误
	num3 := <-intChan
	num4 := <-intChan
	//num5 := <-intChan	//deadlock错误
	fmt.Println("num3=", num3, "num4=", num4)
	fmt.Printf("channel len=%v, cap=%v \n", len(intChan), cap(intChan))

	//map类型演示
	mapChan := make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m2 := make(map[string]string, 20)

	m1["city1"] = "北京"
	m1["city2"] = "上海"

	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"

	//存放
	mapChan <- m1
	mapChan <- m2
	fmt.Printf("mapChan len=%v, cap=%v \n", len(mapChan), cap(mapChan))

	//读取
	fmt.Printf("mapChan m1=%v \n", <-mapChan)
	fmt.Printf("mapChan m2=%v \n", <-mapChan)
	fmt.Printf("mapChan len=%v, cap=%v \n", len(mapChan), cap(mapChan))

	//结构体演示
	catChan := make(chan Cat, 10)
	catChan <- Cat{Name: "tom", Age: 10}
	catChan <- Cat{Name: "jerry", Age: 20}

	//取出
	fmt.Printf("catChan cat1=%v \n", <-catChan)
	fmt.Printf("catChan cat2=%v \n", <-catChan)
	fmt.Printf("catChan len=%v, cap=%v \n", len(catChan), cap(catChan))

	fmt.Println("OK")
}
