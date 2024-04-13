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

	//加一个表明属于哪个用户的连接
	Userid int
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
		if err == model.ERROR_USER_NOTEEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
		//我们先测试成功，再返回具体的错误信息
	} else {
		loginResMes.Code = 200
		//把登陆成功的用户放入userMgr中
		//将登陆成功的用户的userid赋给this
		this.Userid = loginMes.Userid
		userMgr.AddOnlineuser(this)

		this.Notifyotheronlineusers(loginMes.Userid)
		//将当前在线用户的id放入到loginResMes.UserIds
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.Userids = append(loginResMes.Userids, id)
		}

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

func (this *Userprocess) Serverprocessregister(mes *message.Message) (err error) {

	//先从mes中取出mes.Data,直接反序列化成registerMes

	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("JSON.unmarshal fail err =", err)
		return
	}

	//先申明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMestype

	//申明一个loginresmes,并完成赋值
	var registerResMes message.RegisterResMes

	//现在需要到数据库中进行注册
	//1.使用modle.Myuserdao到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_NOTEEXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 200
		}
	} else {
		registerResMes.Code = 200
	}

	//将loginresmes序列化
	data, err := json.Marshal(registerResMes)
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

// 编写一个通知所有人在线的一个方法
// userid通知其他的在线用户，我上线了
func (this *Userprocess) Notifyotheronlineusers(userid int) {

	//遍历UserMgr切片，然后逐个发送Notifyuserstatusmes，上线消息
	for id, up := range userMgr.onlineUsers {

		//过滤自己
		if id == userid {
			continue
		}

		//开始通知，单独写一个方法
		up.Notify(userid)
	}
}

func (this *Userprocess) Notify(userid int) {

	//组装Notifyuserstatusmes
	var mes message.Message
	mes.Type = message.NotifyuserstatusmesType

	var notifywuerstatusmes message.Notifyuserstatusmes
	notifywuerstatusmes.Userid = userid
	notifywuerstatusmes.Status = message.Useronline

	//将notifywuerstatusmes序列化
	data, err := json.Marshal(notifywuerstatusmes)
	if err != nil {
		fmt.Println("json.marshal notify err=", err)
		return
	}

	//将序列化后的功能赋值给mes.data
	mes.Data = string(data)

	//对mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json mes err = ", err)
		return
	}

	//发送，创建transger实例，发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.Writepkg(data)
	if err != nil {
		fmt.Println("notify err = ", err)
		return
	}
}
