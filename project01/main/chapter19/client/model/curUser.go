package model

import (
	"go_code/project01/main/chapter19/common/message"
	"net"
)

//CurUser 因为在客户端我们很多地方会使用到CurUser，我们将其作为全局的
type CurUser struct {
	Conn net.Conn
	message.User
}
