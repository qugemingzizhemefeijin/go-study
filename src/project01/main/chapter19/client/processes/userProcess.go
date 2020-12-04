package processes

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_code/project01/main/chapter19/client/utils"
	"go_code/project01/main/chapter19/common/message"
	"net"
	"os"
)

//UserProcess ...
type UserProcess struct {
	//暂时不需要字段
}

//Register ...
func (userProcess *UserProcess) Register(userID int, userPwd, userName string) (err error) {
	//1. 链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2. 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3. 创建一个RegisterMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserID = userID
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4. 将registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5. 把data赋给mes.Data字段
	mes.Data = string(data)

	//6. 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	fmt.Printf("客户端，发送注册消息的长度=%d 内容=%v\n", len(data), string(data))

	//7. 到这个时候，date就是我们要发送的消息
	//7.1 先把 data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	//这里还需要处理服务器端返回的消息
	mes, err = tf.ReadPkg() //mes 就是
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分分序列化成 RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail=", err)
		return
	}

	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请登录")
	} else {
		fmt.Println(registerResMes.Error)
		err = errors.New(registerResMes.Error)
	}
	os.Exit(0)
	return
}

//Login 写一个函数，完成登录校验
func (userProcess *UserProcess) Login(userID int, userPwd string) (err error) {
	//下一步就要开始订协议...
	//fmt.Printf("userID=%v, userPwd=%v \n", userID, userPwd)

	//return nil

	//1. 链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2. 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3. 创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd

	//4. 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5. 把data赋给mes.Data字段
	mes.Data = string(data)

	//6. 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	fmt.Printf("客户端，发送消息的长度=%d 内容=%v\n", len(data), string(data))

	//7. 到这个时候，date就是我们要发送的消息
	//7.1 先把 data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	//这里还需要处理服务器端返回的消息
	mes, err = tf.ReadPkg() //mes 就是
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分分序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail=", err)
		return
	}

	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserID = userID
		CurUser.UserStatus = message.UserOnline
		//fmt.Println("登录成功")

		//可以显示当前在线用户列表，遍历loginResMes.UserIds
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserIds {
			//如果我们要求不显示自己在线，下面我们增加一个代码
			if v == userID {
				continue
			}
			fmt.Printf("用户ID:%d \n", v)
			//完成 客户端的onlineUsers 初始化
			onlineUsers[v] = &message.User{
				UserID:     v,
				UserStatus: message.UserOnline,
			}
		}
		fmt.Printf("\n\n")

		//此处还要显示用户的离线消息
		offlineMessage := loginResMes.OfflineSmsMes
		if offlineMessage != nil && len(offlineMessage) > 0 {
			fmt.Println("你当前的离线消息列表如下：")
			for _, mess := range offlineMessage {
				fmt.Printf("用户 %d 说：%s \n", mess.FromUserID, mess.Content)
			}
			fmt.Println("==========================")
		}

		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端.
		go ProcessServerMes(conn)

		//1. 显示我们登录成功后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
		err = errors.New(loginResMes.Error)
	}

	return
}
