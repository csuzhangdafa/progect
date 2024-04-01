package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 只有首字母大写才可以导出
func ReadPkg(conn net.Conn) (mes message.Message, err error) {

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