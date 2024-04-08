package process

import (
	"chatroom/client/utils"
	"fmt"
	"net"
	"os"
)

//显示登录成功的菜单，二级菜单

func Showmenu() {
	fmt.Println("------------恭喜xxx登陆成功------------")
	fmt.Println("------------1.显示用户在线列表------------")
	fmt.Println("------------2.发送消息------------")
	fmt.Println("------------3.消息列表------------")
	fmt.Println("------------4.退出系统------------")
	fmt.Println("------------请选择(1-4):------------")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示用户在线列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入选项不对>>>>>")
	}
}

// 和服务器端保持通讯
func ProcesserMes(Conn net.Conn) {
	tf := &utils.Transfer{
		Conn: Conn,
	}
	for {
		fmt.Println("客户端正在等待服务器发送消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.Readpkg err =", err)
			return
		}

		//读取到数据下一步处理
		fmt.Printf("mes=%v", mes)
	}
}
