package main

import (
	"fmt"
	"strconv"
)

func main() {
	//使用内置函数close可以关闭channel，当channel关闭之后，就不能再向channel写数据了，但是扔人可以从该channel读取数据
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	//关闭管道
	close(intChan)

	//这时不能再写入了
	//intChan <- 100	//panic: send on closed channel

	//当管道关闭后，读取是没问题的。如果所有的值都被读取后，n1将返回零值，ok=false，意思是没有任何可读取的信息了
	n1, ok := <-intChan
	fmt.Printf("n1=%d, ok=%v \n", n1, ok)
	n1, ok = <-intChan
	fmt.Printf("n1=%d, ok=%v \n", n1, ok)
	n1, ok = <-intChan //这里ok返回false,n1=0
	fmt.Printf("n1=%d, ok=%v \n", n1, ok)

	//channel的遍历
	//channel支持for-range方式进行遍历，注意两个细节：
	//1.在遍历时，如果channel没有关闭，则会出现deallock的错误
	//2.在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完成后，就会退出遍历。

	strChan := make(chan string, 100)
	for i := 0; i < 100; i++ {
		strChan <- "我爱你-" + strconv.Itoa(i*2)
	}

	//遍历 之前必须要close一下关闭，否则会报deallock错误，但是数据确实是全部遍历出来了
	close(strChan) //如果关闭管道，就可以正常的输出了
	for v := range strChan {
		fmt.Printf("v= %v\n", v)
	}

	fmt.Println("OK")
}
