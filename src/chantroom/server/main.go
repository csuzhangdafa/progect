package main

import (
	"fmt"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	//延时关闭
	defer conn.Close()

	//循环读取客户信息
	for {
		buf := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.read err=", err)
			return
		}
		fmt.Println("读取到的buf=", buf[:4])
	}
}

func main() {
	fmt.Println("服务器在8889端口监听")
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
