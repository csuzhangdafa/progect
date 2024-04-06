package main

import (
	"chatroom/server/processer"
	"fmt"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	//延时关闭
	defer conn.Close()
	//这里调用总控，创建一个Processor结构体
	processer := &processer.Processer{
		Conn: conn,
	}

	err := processer.Process02()
	if err != nil {
		fmt.Println("携程错误", err)
		return
	}
}

func main() {
	fmt.Println("（服务器改进新的结构）服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listenerr=", err)
		return
	}

	//监听成功，等待客户端链接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen,accept err=", err)
		}

		//连接成功启动一个协程和客户端保持通讯
		go process(conn)
	}
}
