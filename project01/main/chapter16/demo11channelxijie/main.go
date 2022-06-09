package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello, World")
	}
}

func test() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误了,err", err)
		}
	}()
	//定义一个map
	var myMap map[int]string
	myMap[0] = "golang" //这里error，因为需要make一下
}

func main() {
	//channel使用细节和注意事项：
	//1. channel可以声明为只读，或者只写性质（默认情况下是双向的）
	//var chan1 chan int //可读可写

	//var chan2 chan<- int //声明只写
	//chan2 = make(chan int, 3)
	//chan2 <- 20
	//num := <-chan2	//这里会报错的

	//var chan3 <-chan int //声明为只读
	//fmt.Println("chan3=", chan3)
	//num2 := <-chan3
	//chan3 <- 30 //错误
	//fmt.Println("num2=", num2)

	//2. channel只读和只写的最佳实践案例 (一般就是用于方法的形参上来控制管道的输入还是输出，防止误操作)

	//3. 使用select可以解决从管道取数据的阻塞问题

	// intChan := make(chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	intChan <- i
	// }

	// stringChan := make(chan string, 5)
	// for i := 0; i < 5; i++ {
	// 	stringChan <- "hello" + fmt.Sprintf("%d", i)
	// }

	//传统的方法，在遍历管道时，如果不关闭会阻塞而导致死锁

	//问题，在实际开发中，有可能我们不好确定什么时候关闭该管道，就可以使用select方式来解决
	// lable:
	// 	for {
	// 		select {
	// 		//注意：这里如果intChan一直没有关闭，也不会一直阻塞而死锁
	// 		//会自动到下一个case匹配
	// 		case v := <-intChan:
	// 			fmt.Printf("从intChan读取的数据%d\n", v)
	// 			time.Sleep(time.Second)
	// 		case v := <-stringChan:
	// 			fmt.Printf("从stringChan读取的数据%v\n", v)
	// 			time.Sleep(time.Second)
	// 		default:
	// 			fmt.Printf("都取不到了，不玩了，程序员可以加入自己的业务逻辑\n")
	// 			break lable
	// 		}
	// 	}

	//4. goroutine中使用recover，解决协程中出现panic，导致程序崩溃问题。（如果我们起了一个协程，但是这个协程出现了panic，如果我们没有
	//   捕获这个panic，就会造成整个程序崩溃，这是我们可以在goroutine中使用recover来捕获panic，进行处理，这样即使这个协程发生了问题
	//   也不影响主线程，可以继续执行）

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok i=", i)
		time.Sleep(time.Second)
	}

	fmt.Println("exit success")
}
