package processes

import (
	"encoding/json"
	"fmt"
	"go_code/project01/main/chapter19/client/utils"
	"go_code/project01/main/chapter19/common/message"
	"net"
	"os"
)

//ShowMenu 显示登录成功后的界面...
func ShowMenu() {
	fmt.Println("---------------恭喜xxxx登录成功----------------")
	fmt.Println("1. 显示在线用户列表")
	fmt.Println("2. 发送群聊消息")
	fmt.Println("3. 发送私聊消息")
	fmt.Println("4. 信息列表")
	fmt.Println("5. 退出系统")
	fmt.Println("请选择(1-5):")

	var key int
	var content string
	var friendUserID int

	//因为，我们总会使用到SmsProcess实例，因此我们将其定义在switch外部
	sp := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表~")
		outputOnlineUser()
	case 2: //群聊
		fmt.Println("你想对大家说点什么~")
		fmt.Scanf("%s\n", &content)
		sp.SendGroupSms(content)
	case 3: //私聊
		fmt.Println("请输入好友ID:")
		fmt.Scanf("%d\n", &friendUserID)
		fmt.Println("请输入内容:")
		fmt.Scanf("%s\n", &content)
		sp.SendToFriendSms(friendUserID, content)
	case 4:
		fmt.Println("信息列表~")
	case 5:
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确~")
	}
}

//ProcessServerMes 和服务器端保持通讯
func ProcessServerMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//如果读取到消息，有事下一步处理消息
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线/离线了
			//1. 取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			fmt.Println("notifyUserStatusMes=", notifyUserStatusMes)
			//2. 把这个用户的信息，状态保存到客户端map中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: // 接收到服务器的群发消息
			outputGroupMes(&mes)
		case message.SmsOneMesType: //接收到服务器的私聊消息
			outputFriendMes(&mes)
		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}
		//fmt.Printf("mes=%v\n", mes)
	}
}
