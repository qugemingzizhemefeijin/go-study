package processes

import (
	"encoding/json"
	"fmt"
	"go_code/project01/main/chapter19/common/message"
)

//显示其它用户发送的消息
func outputGroupMes(mes *message.Message) {
	//1.反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	//fmt.Println("mes=", mes)

	//显示信息
	info := fmt.Sprintf("用户ID\t%d 对大家说:\t%s\n\n", smsMes.UserID, smsMes.Content)
	fmt.Println(info)
}

//显示其它用户发送的私聊消息
func outputFriendMes(mes *message.Message) {
	//1.反序列化
	var smsOneMes message.SmsOneMes
	err := json.Unmarshal([]byte(mes.Data), &smsOneMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	//显示信息
	info := fmt.Sprintf("用户ID\t%d 私聊:\t%s\n\n", smsOneMes.UserID, smsOneMes.Content)
	fmt.Println(info)
}
