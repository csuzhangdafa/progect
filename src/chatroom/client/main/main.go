package main

import (
	"chatroom/client/process"
	"fmt"
)

// 用户id与用户密码
var userid int
var userpwd string
var username string

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
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userid)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userpwd)
			//loop = false
			//完成登录，创建一个UserProcess实例
			up := &process.Userprocess{}
			up.Login(userid, userpwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userid)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userpwd)
			fmt.Println("请输入用户名字")
			fmt.Scanf("%s\n", &username)
			//loop = false
			//完成注册，创建一个UserProcess实例
			//up := &process.Userprocess{}

		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你输入有误，请重新输入1-3")
		}
	}
}
