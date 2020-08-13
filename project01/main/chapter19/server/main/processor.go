package main

import (
	"fmt"
	"go_code/project01/main/chapter19/common/message"
	"go_code/project01/main/chapter19/server/processes"
	"go_code/project01/main/chapter19/server/utils"
	"io"
	"net"
)

//Processor 先创建一个Processor 的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes 函数
//功能：根据客户端发送的消息种类不同，决定调用哪个函数来处理
func (processor *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑
		//创建一个UserProcess实例
		up := &processes.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册的逻辑
		up := &processes.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//处理用户群发，创建SmsProcess实例完成转发群聊消息。
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	case message.SmsOneMesType:
		//处理私聊消息，完成消息中转
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendMesToUser(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (processor *Processor) processHandler() (err error) {
	defer func() {
		//当次函数退出，则代表用户链接断开了，需要清理用户的在线状态以及给其他用户发送离线消息
		up := &processes.UserProcess{
			Conn: processor.Conn,
		}
		up.NotifyOfflineUser()
	}()
	//循环读取客户端发送的消息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
		//创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: processor.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端 %s 退出，服务器端也退出 \n", processor.Conn.RemoteAddr().String())
			} else {
				fmt.Println("readPkg  err=", err)
			}
			return err
		}
		//fmt.Println("mes=", mes)
		err = processor.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes err=", err)
			return err
		}
	}
}
