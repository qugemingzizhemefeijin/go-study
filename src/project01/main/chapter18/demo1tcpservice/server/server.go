package main

import (
	"fmt"
	"io"
	"net"
)

//全局的客户端连接映射
var clientMap map[int]net.Conn = make(map[int]net.Conn)

func process(conn net.Conn, idx int) {
	//这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn

	//维护conn
	clientMap[idx] = conn

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//服务端在此处等待conn发送信息
		//如果客户端没有write，那么此处协程会一直阻塞
		fmt.Println("服务器在等待客户端发送信息" + conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端" + conn.RemoteAddr().String() + "已退出")
			} else {
				fmt.Println("服务器的conn.Read() err=", err)
			}
			break
		}

		//我收到的信息结果
		str := string(buf[:n])
		fmt.Printf("接收到客户端 %d 的信息：%v \n", idx, str)

		//给另一个客户端发送接收到的信息
		otherConn := clientMap[1]
		if otherConn != nil {
			fmt.Println("给另一个客户端发送消息")
			otherConn.Write([]byte(str))
		}
	}
}

//1.编写一个服务器端程序，在8888端口监听
//2.可以和多个客户端创建链接
//3.链接成功后，客户端可以发送数据，服务器端接受数据，并显示在终端上
//4.先使用telnet来测试，然后编写客户端程序来测试
func main() {
	fmt.Println("服务器开始监听...")
	//1. tcp表示使用网络协议是tcp
	//2. 0.0.0.0:8888表示在本地监听8888端口
	//可以127.0.0.1:8888 支持ipv4
	//可以0.0.0.0:8888 支持ipv4和ipv6
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	fmt.Printf("listen=%v \n", listen)

	//延时关闭listen
	defer listen.Close()

	idx := 0
	//循环等待客户端来连接我
	for {
		fmt.Println("等待客户端来连接...")
		//Accept等待并返回下一个连接到该接口的连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
			continue
		}
		fmt.Printf("Accept() success conn=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		//这里准备起一个协程，为客户端服务
		go process(conn, idx)
		idx++
	}

	//fmt.Println("exit success")
}
