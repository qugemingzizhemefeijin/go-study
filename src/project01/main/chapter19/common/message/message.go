package message

const (
	//LoginMesType ...
	LoginMesType = "LoginMes"
	//LoginResMesType ...
	LoginResMesType = "LoginResMes"
	//RegisterMesType ...
	RegisterMesType = "RegisterMes"
	//RegisterResMesType ...
	RegisterResMesType = "RegisterRes"
	//NotifyUserStatusMesType ...
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	//SmsMesType ...
	SmsMesType = "SmsMes"
	//SmsOneMesType ...
	SmsOneMesType = "SmsOneMes"
)

//这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

//Message 消息体
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

//定义两个消息...

//LoginMes ...
type LoginMes struct {
	UserID   int    `json:"userId"`   //用户ID
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名称
}

//LoginResMes ...
type LoginResMes struct {
	Code          int             `json:"code"`          //状态码 500表示该用户未注册,200表示成功
	UserIds       []int           `json:"userIds"`       //保存用户ID的切片
	OfflineSmsMes []OfflineSmsMes `json:"offlineSmsMes"` //保存用户的离线消息
	Error         string          `json:"err"`           //错误信息
}

//RegisterMes ...
type RegisterMes struct {
	User User `json:"user"` //类型是message结构体
}

//RegisterResMes ...
type RegisterResMes struct {
	Code  int    `json:"code"`  //返回状态码 400 表示该用户已经占用 200表示注册成功
	Error string `json:"error"` //错误信息
}

//NotifyUserStatusMes 为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserID int `json:"userId"`
	Status int `json:"status"` //用户状态
}

//SmsMes 增加一个SmsMes 结构体
type SmsMes struct {
	User           //匿名结构体，继承
	Content string `json:"content"`
}

//SmsOneMes 私聊
type SmsOneMes struct {
	User            //匿名结构体，继承
	ToUserID int    `json:"toUserId"` //对方ID
	Content  string `json:"content"`
}

//OfflineSmsMes 增加离线消息结构体
type OfflineSmsMes struct {
	FromUserID int    `json:"fromUserId"` //来自哪个用户
	Content    string `json:"content"`
}
