package message

import (
	"chatroom/server/model"
)

const (
	LoginMestype            = "LoginMes"
	LoginResMestype         = "LoginResMes"
	RegisterMestype         = "RegisterMes"
	RegisterResMestype      = "RegisterResMes"
	NotifyuserstatusmesType = "Notifyuserstatusmes"
)

// 定义几个用户状态的常量
const (
	Useronline = iota
	Useroffline
	Userbusystatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 定义两个消息，两个返回类型
type LoginMes struct {
	Userid   int    `json:"userid"`  //用户ID
	Userpwd  string `json:"userpwd"` //用户密码
	Username string `json:"username"`
}

type LoginResMes struct {
	Code    int    `json:"code"` //返回状态码
	Userids []int  //增加字段保存用户id的切片
	Error   string `json:"error"` //返回错误信息
}

type RegisterMes struct {
	//注册登录消息类型
	User model.User //类型就是User结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //400,占用，200表示注册成功
	Error string `json:"error"` //返回错误信息
}

//为了配合服务器端推送用户状态变化的消息
//服务器主动推送，并不需要响应

type Notifyuserstatusmes struct {
	Userid int `json:"userid"` //用户id
	Status int `json:"status"` //用户状态
}
