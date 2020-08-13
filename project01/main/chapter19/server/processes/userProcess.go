package processes

import (
	"encoding/json"
	"fmt"
	"go_code/project01/main/chapter19/common/message"
	"go_code/project01/main/chapter19/server/model"
	"go_code/project01/main/chapter19/server/utils"
	"net"
)

//UserProcess ...
type UserProcess struct {
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户的
	UserID int
}

//NotifyOfflineUser 通知用户离线
func (userProcess *UserProcess) NotifyOfflineUser() {
	//1.首先后去离线的用户
	offlineUserID := userMgr.GetUserIDByConn(userProcess.Conn)
	fmt.Println("获取离线用户ID=", offlineUserID)
	if offlineUserID == 0 {
		return
	}
	//2.将其从在线列表中删除掉
	userMgr.DeleteOnlineUser(offlineUserID)
	//3.循环离线推送消息
	for id, up := range userMgr.onlineUsers {
		fmt.Printf("发送离线数据%d 给 %d \n", offlineUserID, id)
		//开始通知 【单独的写一个方法】
		up.NotifyMeOnline(offlineUserID, message.UserOffline)
	}
}

//NotifyOthersOnlineUser 这里我们编写通知所有在线的用户的方法
//要通知其他的在线用户，我上线了
func (userProcess *UserProcess) NotifyOthersOnlineUser() {
	//需要遍历userMgr.onlineUsers ，然后一个一个发送NotifyUserStatusMes消息
	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if userProcess.UserID == id {
			continue
		}
		fmt.Printf("发送上线数据%d 给 %d \n", userProcess.UserID, id)
		//开始通知 【单独的写一个方法】
		up.NotifyMeOnline(userProcess.UserID, message.UserOnline)
	}
}

//NotifyMeOnline 给每个用户发送指定用户上线了
//userID 状态变更人
//status 状态
func (userProcess *UserProcess) NotifyMeOnline(userID int, status int) {
	//组装我们的NoticyUserStatusRes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = userID
	notifyUserStatusMes.Status = status

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//再次对mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送消息，创建我们Transfer实例，发送
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
	fmt.Printf("发送登录数据%d 给 %d 成功 \n", userID, userProcess.UserID)
}

//ServerProcessRegister 专门处理注册请求
func (userProcess *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1. 先从mes 中取出 mes.Data，并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	//反序列化json
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//再申明一个RegisterResMes
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库中去完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 500
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "服务器内部错误..."
		}
	} else {
		fmt.Println("user=", registerMes.User)
		registerResMes.Code = 200
		registerResMes.Error = "注册成功"
	}

	//进行反序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	//4. 将data赋值给resMes
	resMes.Data = string(data)

	//5. 将resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	//6. 发送data，我们将其封装到writePkg函数中
	//因为使用了分层模式，我们先创建Transfer实例，然后来写入
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//ServerProcessLogin 专门处理登录请求
func (userProcess *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码...
	//1. 先从mes 中取出 mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	//反序列化成json
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个 LoginResMes
	var loginResMes message.LoginResMes

	//我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao 到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}

		//这里我们先测试成功，然后我们可以根据返回具体错误信息
	} else {
		loginResMes.Code = 200
		loginResMes.Error = "登录成功"

		//将登录成功用户ID赋值一下
		userProcess.UserID = loginMes.UserID
		//这里因为用户登录成功，我们就把该登录成功的用户放入到userMgr中
		userMgr.AddOnlineUser(userProcess)
		//通知其他的用户我上线了
		userProcess.NotifyOthersOnlineUser()
		//将当前在线用户的ID 放入到loginResMeg.UserIds切片中
		//遍历 userMgr.onlineUsers
		for id := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		//此处还需要放置用户的离线消息
		offlineMessage, err := model.MyUserDao.GetUserAllOfflineMessage(loginMes.UserID)
		if err == nil {
			loginResMes.OfflineSmsMes = offlineMessage
		}

		fmt.Println(user, "登录成功")
	}

	//如果用户ID=100，密码=123456，认为合法，否则不合法
	// if loginMes.UserID == 100 && loginMes.UserPwd == "123456" {
	// 	//合法
	// 	loginResMes.Code = 200
	// 	loginResMes.Error = "登录成功"
	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500 //500 状态码，边是该用户不存在
	// 	loginResMes.Error = "该用户不存在，请注册再使用"
	// }

	//3. 将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	//4. 将data赋值给resMes
	resMes.Data = string(data)

	//5. 将resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail=", err)
		return
	}

	//6. 发送data，我们将其封装到writePkg函数中
	//因为使用了分层模式，我们先创建Transfer实例，然后来写入
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}
