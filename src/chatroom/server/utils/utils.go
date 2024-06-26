package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 将这些方法关联到结构体中，使用oop思想
type Transfer struct {

	//分析应该存在的字段
	Conn net.Conn
	Buf  [8096]byte
}

// 只有首字母大写才可以导出
func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")

	//conn.read在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn则，就不会被阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	//根据buf[:4]转成一个uin32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn.read fail err=", err)
		return
	}

	//把pakLen反序列化成message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.unmarsha err=", err)
		return
	}
	return
}

func (this *Transfer) Writepkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkglen uint32
	pkglen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkglen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buyes) fail", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.write(data) fail", err)
		return
	}
	return
}
