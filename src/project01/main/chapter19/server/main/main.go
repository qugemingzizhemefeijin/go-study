package main

import (
	"fmt"
	"go_code/project01/main/chapter19/server/model"
	"net"
	"time"
)

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里也需要延时关闭conn
	defer conn.Close()

	//这里调用总控，需要创建一个Processor
	processor := &Processor{
		Conn: conn,
	}
	err := processor.processHandler()
	if err != nil {
		fmt.Println("客户端和服务器端通讯协程错误=", err)
		return
	}
}

//这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	//这里的pool本身就是一个全局的变量
	//这里需要注意一个初始化顺序问题
	//initPool，在 initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//当服务器启动时，我们就去初始化redis的连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	//go build -o server.exe go_code/project01/main/chapter19/server/main
	fmt.Println("服务器[新的结构]在8889端口监控...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		} else {
			//一旦连接成功，则启动一个协程和客户端保持通讯
			go process(conn)
		}
	}
	//fmt.Println("exit success")
}
