package processer

import (
	"chatroom/common/message"
	"chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

// 先创建一个Processor结构体
type Processer struct {
	Conn net.Conn
}

// 根据客户端发送的消息种类不同，决定用哪个函数来处理
func (this *Processer) Serverprocessmes(mes *message.Message) (err error) {

	//看看从客户端是否能接收到群发的消息
	fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMestype:
		//处理登录
		//创建一个Userprocesser实例
		up := &process.Userprocess{
			Conn: this.Conn,
		}
		err = up.Serverprocesslogin(mes)
	case message.RegisterMestype:
		//处理注册
		up := &process.Userprocess{
			Conn: this.Conn,
		}
		err = up.Serverprocessregister(mes) //type:data
	case message.Smsmestype:
		//创建一个smsprocess转发群聊消息的实例
		smsprocess := &process.Smeprocess{}
		smsprocess.Sendgroupmes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processer) Process02() (err error) {
	//循环读取客户信息
	for {

		//创建transfer完成读包的任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也正常退出..")
				return err
			} else {
				fmt.Println("read err =", err)
				return err
			}
		}
		fmt.Println("mes=", mes)
		err = this.Serverprocessmes(&mes)
		if err != nil {
			return err
		}
	}
	return err
}
