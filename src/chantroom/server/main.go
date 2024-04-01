package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")

	//conn.read在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn则，就不会被阻塞
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	//根据buf[:4]转成一个uin32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen读取消息内容
	n, err = conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn.read fail err=", err)
		return
	}

	//把pakLen反序列化成message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.unmarsha err=", err)
		return
	}
	return
}

func writepkg(conn net.Conn, data []byte) (err error) {

	//先发送一个长度给对方
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buyes) fail", err)
		return
	}

	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.write(data) fail", err)
		return
	}
	return
}

// 编写一个serverprocesslogin函数，专门处理登录请求
func serverprocesslogin(conn net.Conn, mes *message.Message) (err error) {

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

	//如果id =100,密码=123456，认为合法，否则不合法
	if loginMes.Userid == 100 && loginMes.Userpwd == "123456" {
		//合法
		loginResMes.Code = 200 //200成功
	} else {
		//不合法
		loginResMes.Code = 500 //500不成功
		loginResMes.Error = "用户不存在，请注册再使用..."
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
	err = writepkg(conn, data)
	return
}

// 编写一个serverprocessmes函数
// 功能：根据客户端发送的消息不同，决定调用哪个函数来处理
func serverprocessmes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMestype:
		//处理登录
		err = serverprocesslogin(conn, mes)
	case message.RegisterMestype:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

// 处理和客户端的通讯
func process(conn net.Conn) {
	//延时关闭
	defer conn.Close()

	//循环读取客户信息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也正常退出..")
				return
			} else {
				fmt.Println("read err =", err)
				return
			}
		}
		fmt.Println("mes=", mes)

		err = serverprocesslogin(conn, &mes)
		if err != nil {
			return
		}
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
