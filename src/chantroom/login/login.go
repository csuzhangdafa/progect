package login

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 登录函数
func Login(userid int, userpwd string) (err error) {

	//定协议
	//fmt.Println("userid = %d userpwd = %s\n",userid,userpwd)

	//return nil
	//连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	//延时关闭
	defer conn.Close()

	//准备通过conn发消息给服务器
	var mes message.Message
	mes.Type = message.LoginMestype

	//创建一个logininMes结构体
	var loginMes message.LoginMes
	loginMes.Userid = userid
	loginMes.Userpwd = userpwd

	//将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//把data赋值给mes.Data字段
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}

	//data就是我们要发送的数据
	//先把data的长度发给服务器
	//获取到data的长度->转成一个表示长度的byte切片
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buyes) fail", err)
		return
	}

	fmt.Println("客户端发送消息的长度=%d", len(data))

	return
}
