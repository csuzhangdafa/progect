package main

import (
	"chatroom/login"
	"fmt"
)

// 用户id与用户密码
var userid int
var userpwd string

func main() {

	var key int
	var loop = true
	for loop {
		fmt.Println("------------欢迎登陆-----------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天系统")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你输入有误，请重新输入1-3")
		}
	}

	//根据用户的输入，选择新的提示信息
	if key == 1 {
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userid)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userpwd)

		err := login.Login(userid, userpwd)
		if err != nil {
			fmt.Println("登陆失败")
			fmt.Println(err)
		} else {
			fmt.Println("登陆成功")
		}
		//登陆函数写到另一个文件
	} else if key == 2 {
		fmt.Println("用户注册")
	}
}
