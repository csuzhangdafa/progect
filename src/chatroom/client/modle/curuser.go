package modle

import (
	"chatroom/common/message"
	"net"
)

// 在客户端都可能用到，所以我们把它作为一个全局变量
type Curuser struct {
	Conn net.Conn
	message.User
}
