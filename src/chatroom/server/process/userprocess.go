package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type Userprocess struct {
	Conn net.Conn
}

// 编写一个serverprocesslogin方法，专门处理用户登录请求
func (this *Userprocess) Serverprocesslogin(mes *message.Message) (err error) {

	//先从mes中取出mes.Data,直接反序列化成loginmes

	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("JSON.unmarshal fail err =", err)
		return
	}

	//先申明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMestype

	//申明一个loginresmes,并完成赋值
	var loginResMes message.LoginResMes

	//现在需要到数据库中进行验证
	//1.使用modle.Myuserdao到redis去验证
	user, err := model.MyUserDao.Login(loginMes.Userid, loginMes.Userpwd)

	if err != nil {
		fmt.Println("开始比较-----")
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在"
		//我们先测试成功，再返回具体的错误信息
	} else {
		fmt.Println("开始比较")
		loginResMes.Code = 200
		fmt.Println(user, "登陆成功")
	}

	//将loginresmes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshal fail", err)
		return
	}

	//将data赋值给resmes
	resMes.Data = string(data)

	//对resmes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal fail", err)
		return
	}

	//发送data数据，我们将其封装成write函数
	//因为使用了分层模式(mvc)我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.Writepkg(data)
	return
}
