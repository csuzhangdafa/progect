package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type Smeprocess struct {
	//暂时不需要字段
}

// 转发消息的方法
func (this *Smeprocess) Sendgroupmes(mes *message.Message) {
	//遍历服务器的onlineusers map[int]*userprocess
	//将消息转发
	//取出mes的内容smsmes
	var smsmes message.Smsmes
	err := json.Unmarshal([]byte(mes.Data), &smsmes)
	if err != nil {
		fmt.Println("json,Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {

		//还需要过滤自己，不要发给自己
		if id == smsmes.Userid {
			continue
		}
		this.sendmestoeach(data, up.Conn)
	}
}

func (this *Smeprocess) sendmestoeach(data []byte, conn net.Conn) {
	//创建一个transfer实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.Writepkg(data)
	if err != nil {
		fmt.Println("转发消息失败err=", err)
	}
}
