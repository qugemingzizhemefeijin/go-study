package processes

import (
	"encoding/json"
	"fmt"
	"go_code/project01/main/chapter19/common/message"
	"go_code/project01/main/chapter19/server/model"
	"go_code/project01/main/chapter19/server/utils"
	"net"
)

//SmsProcess ...
type SmsProcess struct {
}

//SendMesToUser 转发私聊消息
func (sp *SmsProcess) SendMesToUser(msg *message.Message) {
	//先解析出SmsOneMes 实例
	var smsOneMes message.SmsOneMes
	err := json.Unmarshal([]byte(msg.Data), &smsOneMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//fmt.Println("msg=", msg)

	//这里同时也要查看用户是否在线
	up, err := userMgr.GetOnlineUserByID(smsOneMes.ToUserID)
	if err != nil { //后面可以发送到redis保存，不过此处还得查询一下是否是用户未上线还是不存在此用户
		//fmt.Printf("用户 %d 当前不在线或不存在，无法中转消息 err=%v \n", smsOneMes.ToUserID, err)
		//此处异常是因为用户可能不存在或者不在线，需要先查询一下redis
		b, err := model.MyUserDao.ExistsUserByID(smsOneMes.ToUserID)
		if err != nil {
			fmt.Printf("查看用户 %d 是否存在发生错误，err=%v \n", smsOneMes.ToUserID, err)
			return
		}
		if b {
			offlineMessage := &message.OfflineSmsMes{
				FromUserID: smsOneMes.UserID,
				Content:    smsOneMes.Content,
			}
			//保存到离线消息列表中
			model.MyUserDao.SaveUserOfflineMessage(smsOneMes.ToUserID, offlineMessage)
		} else {
			fmt.Printf("用户 %d 不存在 \n", smsOneMes.ToUserID)
		}
		return
	}

	//将msg序列化，以便转发给别人
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//fmt.Println("消息中转给用户", up.UserID)

	//将消息转发给被人
	sp.SendMesToEachOnlineUser(data, up.Conn)
}

//SendGroupMes 转发消息
func (sp *SmsProcess) SendGroupMes(msg *message.Message) {
	//先解析出SmsMes实例
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(msg.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//将msg序列化，以便转发给别人
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//fmt.Println("smsMes=", smsMes)

	//如果要群聊发送给所有人，这里可以稍微改动一下
	users, err := model.MyUserDao.GetAllUsers()
	if err != nil {
		fmt.Println("GetAllUsers() err=", err)
		return
	}

	for _, user := range users {
		if smsMes.UserID != user.UserID {
			up, err := userMgr.GetOnlineUserByID(user.UserID)
			if err != nil { //用户不在线，则发送离线消息
				offlineMessage := &message.OfflineSmsMes{
					FromUserID: smsMes.UserID,
					Content:    smsMes.Content,
				}
				//保存到离线消息列表中
				model.MyUserDao.SaveUserOfflineMessage(user.UserID, offlineMessage)
			} else {
				sp.SendMesToEachOnlineUser(data, up.Conn)
			}
		}
	}
	//遍历服务器端的onlineUsers map[int]*UserProcess，将消息转发出去
	// for id, up := range userMgr.onlineUsers {
	// 	//需要过滤掉自己
	// 	if smsMes.UserID != id {
	// 		sp.SendMesToEachOnlineUser(data, up.Conn)
	// 	}
	// }
}

//SendMesToEachOnlineUser 转发消息给指定的链接
func (sp *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("Send To EachOnline User err=", err)
	}
	return
}
