package processes

import (
	"fmt"
	"net"
)

//因为UserMgr 实例在服务器端有且只有一个
//因为在很多的地方，都会使用到，因为此，我们将其定义为全局变量

var (
	userMgr *UserMgr
)

//UserMgr ...
type UserMgr struct {
	onlineUsers map[int]*UserProcess
	connUserIDS map[net.Conn]int
}

//完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
		connUserIDS: make(map[net.Conn]int, 1024),
	}
}

//AddOnlineUser 完成对onlineUsers添加
func (userMgr *UserMgr) AddOnlineUser(up *UserProcess) {
	userMgr.onlineUsers[up.UserID] = up
	userMgr.connUserIDS[up.Conn] = up.UserID
}

//DeleteOnlineUser 删除
func (userMgr *UserMgr) DeleteOnlineUser(userID int) {
	up, ok := userMgr.onlineUsers[userID]

	delete(userMgr.onlineUsers, userID)
	if ok {
		delete(userMgr.connUserIDS, up.Conn)
	}
}

//GetUserIDByConn 根据链接获取用户ID
func (userMgr *UserMgr) GetUserIDByConn(conn net.Conn) int {
	return userMgr.connUserIDS[conn]
}

//GetAllOnlineUsers 返回当前所有在线的用户
func (userMgr *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return userMgr.onlineUsers
}

//GetOnlineUserByID 根据ID返回对应的值
func (userMgr *UserMgr) GetOnlineUserByID(userID int) (up *UserProcess, err error) {
	up, ok := userMgr.onlineUsers[userID]
	if !ok { //说明你要查找的这个用户当前不在线
		err = fmt.Errorf("用户%d 不存在", userID)
		return
	}
	return
}
