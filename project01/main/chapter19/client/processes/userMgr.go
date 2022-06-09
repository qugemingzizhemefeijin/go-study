package processes

import (
	"fmt"
	"go_code/project01/main/chapter19/client/model"
	"go_code/project01/main/chapter19/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

//CurUser 全局的CurUser，我们在用户登录成功后，完成对CurUser初始化
var CurUser model.CurUser

//在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for userID := range onlineUsers {
		//如果不显示自己，可以过滤一下
		fmt.Println("用户ID:\t", userID)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//适当的优化
	user, ok := onlineUsers[notifyUserStatusMes.UserID]
	if !ok { //如果没有获取到在线用户信息
		if notifyUserStatusMes.Status == message.UserOffline { //传递过来的状态是离线，则什么都不管
			return
		}
		user = &message.User{
			UserID:     notifyUserStatusMes.UserID,
			UserStatus: notifyUserStatusMes.Status,
		}
		onlineUsers[notifyUserStatusMes.UserID] = user
	} else {
		//如果成功后去到了信息，传递过来的是离线状态，则删除掉
		if notifyUserStatusMes.Status == message.UserOffline {
			delete(onlineUsers, notifyUserStatusMes.UserID)
		} else {
			user.UserStatus = notifyUserStatusMes.Status
		}
	}

	//更新完成后顺带调用一下
	outputOnlineUser()
}

//判断指定用户是否在我的好友在线列表中
func existsUserOnline(userID int) (b bool) {
	//查看用户是否在线
	user, ok := onlineUsers[userID]
	if !ok {
		return
	}
	if user.UserStatus == message.UserOnline {
		b = true
	}
	return
}
