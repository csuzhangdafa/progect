package message

import (
	"chatroom/server/model"
)

const (
	LoginMestype       = "LoginMes"
	LoginResMestype    = "LoginResMes"
	RegisterMestype    = "RegisterMes"
	RegisterResMestype = "RegisterResMes"
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
	Code  int    `json:"code"`  //返回状态码
	Error string `json:"error"` //返回错误信息
}

type RegisterMes struct {
	//注册登录消息类型
	User model.User //类型就是User结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //400,占用，200表示注册成功
	Error string `json:"error"` //返回错误信息
}
