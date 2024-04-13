package process

import (
	"chatroom/common/message"
	"fmt"
)

//客户端要维护的map

var onlineusers map[int]*message.User = make(map[int]*message.User, 10)

// 编写一个方法，处理返回的信息
func Updateuserstatus(notifyuserstatusmes *message.Notifyuserstatusmes) {

	user, ok := onlineusers[notifyuserstatusmes.Userid]
	if !ok {
		//map中没有
		user = &message.User{
			Userid: notifyuserstatusmes.Userid,
		}
	}

	user.Userstatus = notifyuserstatusmes.Status
	onlineusers[notifyuserstatusmes.Userid] = user

	outputonlineuser()
}

// 在客户端显示当前在线的用户
func outputonlineuser() {
	//遍历一遍 onlineusers
	fmt.Println("当前用户列表")
	for id, user := range onlineusers {
		fmt.Println("用户id:\t", id)
		fmt.Println("用户:\t", user)
	}
}
