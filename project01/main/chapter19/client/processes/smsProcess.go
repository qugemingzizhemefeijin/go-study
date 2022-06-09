package processes

import (
	"encoding/json"
	"fmt"
	"go_code/project01/main/chapter19/client/utils"
	"go_code/project01/main/chapter19/common/message"
)

//SmsProcess ...
type SmsProcess struct {
}

//SendToFriendSms 给指定好友发送消息
func (sms *SmsProcess) SendToFriendSms(friendUserID int, content string) (err error) {
	//首先查看此用户是否在线
	// if !existsUserOnline(friendUserID) {
	// 	err = fmt.Errorf("用户%d不在线，无法发送聊天信息", friendUserID)
	// 	fmt.Println(err)
	// 	return
	// }
	//如果在线的话，可以发送私聊消息了
	var mes message.Message
	mes.Type = message.SmsOneMesType

	//创建SmsOneMes实例
	var smsOneMes message.SmsOneMes
	smsOneMes.Content = content
	smsOneMes.ToUserID = friendUserID
	smsOneMes.UserID = CurUser.UserID
	smsOneMes.UserStatus = CurUser.UserStatus

	//将SmsOneMes序列化
	data, err := json.Marshal(smsOneMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err.Error())
		return
	}
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err.Error())
		return
	}

	//调用tf发送消息
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendToFriendSms err=", err.Error())
	}
	return
}

//SendGroupSms 群发
func (sms *SmsProcess) SendGroupSms(content string) (err error) {
	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2. 创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserID = CurUser.UserID
	smsMes.UserStatus = CurUser.UserStatus

	//3.将smsMes序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupSms json.Marshal fail=", err.Error())
		return
	}
	mes.Data = string(data)

	//4. 对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupSms json.Marshal fail=", err.Error())
		return
	}

	//5. 将序列化后的数据发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
	}
	return
}
