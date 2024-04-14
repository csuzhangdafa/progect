package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
)

type Smsprocess struct {
}

// 发送群聊消息
func (this *Smsprocess) sendgroupmes(content string) (err error) {

	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.Smsmestype

	//2.创建一个smsmes实例
	var smsmes message.Smsmes
	smsmes.Content = content
	smsmes.Userid = Curuser.Userid
	smsmes.Userstatus = Curuser.Userstatus

	//序列化smsmes
	data, err := json.Marshal(smsmes)
	if err != nil {
		fmt.Println("sendgroup json.marshal fail =", err.Error())
		return
	}

	mes.Data = string(data)

	//对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendgroup json.marshal fail =", err.Error())
		return
	}

	//将mes发送给服务器
	tf := &utils.Transfer{
		Conn: Curuser.Conn,
	}

	//发送
	err = tf.Writepkg(data)
	if err != nil {
		fmt.Println("sendgroupmes err =", err.Error())
		return
	}
	return
}
