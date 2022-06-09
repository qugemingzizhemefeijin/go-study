package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn) {
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			if err == io.EOF {
				fmt.Println("已退出")
			} else {
				fmt.Println("conn read err=", err)
			}
			break
		}

		//我收到的信息结果
		str := string(buf[:n])
		fmt.Printf("接收到消息 服务器转发 的信息：%v \n", str)
	}
}

//1.编写一个客户端程序，能链接到服务器端的8888端口
//2.客户端可以发送 单行数据，然后就退出
//3.能通过终端输入数据(输入一行发送一行),并发送给服务器端
//4.在终端输入exit，表示退出程序
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn 成功=", conn)
	defer conn.Close()

	//开启协程，接收消息
	go process(conn)

	//功能一：客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入[终端]

	for {
		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n') //line是会包含\n换行符的
		if err != nil {
			fmt.Println("readString err=", err)
			break
		}
		//line = strings.Trim(line, " \r\n")
		fmt.Println("send line=", line)
		if strings.TrimSpace(line) == "exit" {
			break
		}
		//再将line 发送给服务器
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err=", err)
			continue
		}
		fmt.Printf("客户端发送了 %d 个字节的数据\n", n)
	}
	fmt.Println("客户端退出...")
}
