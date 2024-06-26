package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Userprocess struct {
	//暂时不需要字段
}

//关联一个用户登陆的方法

// 登录函数
func (this *Userprocess) Login(userid int, userpwd string) (err error) {

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
	//获取到data的长度转成一个表示长度的byte切片
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buyes) fail", err)
		return
	}

	fmt.Printf("客户端发送的消息长度= %d 内容=%s", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write(data) fail", err)
		return
	}

	//休眠20
	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠20s")

	//处理服务器端返回的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readpkg(conn)err = ", err)
		return
	}

	//将mes的data部分反序列化成loginresmes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println("登陆成功")

		//对curuser初始化
		Curuser.Conn = conn
		Curuser.Userid = userid
		Curuser.Userstatus = message.Useronline

		//显示当前用户在线列表，遍历loginResmes.Userid
		fmt.Println("当前用户列表如下：")
		for _, v := range loginResMes.Userids {

			//不显示自己，加一个代码
			if v == userid {
				continue
			}

			fmt.Println("用户id:\t", v)

			//完成初始化onlineusers
			user := &message.User{
				Userid:     v,
				Userstatus: message.Useronline,
			}
			onlineusers[v] = user
		}
		fmt.Println("\n\n")

		//我们还要在客户端启动一个协程
		//该协程保持和服务器的通讯，如果有数据推送给客户端，则接受并显示
		go ProcesserMes(conn)

		//1.循环显示登陆成功的菜单
		for {
			Showmenu()
		}

	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

// 请求注册的方法
func (this *Userprocess) Register(userid int, userpwd string, username string) (err error) {

	//1.先连接
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
	mes.Type = message.RegisterMestype

	//创建一个logininMes结构体
	var registerMes message.RegisterMes
	registerMes.User.Userid = userid
	registerMes.User.Userpwd = userpwd
	registerMes.User.Username = username

	//将registerMes序列化
	data, err := json.Marshal(registerMes)
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

	//创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务器端
	err = tf.Writepkg(data)
	if err != nil {
		fmt.Println("注册发送信息出错,err = ", err)
		return
	}

	//读取服务端返回的消息
	mes, err = tf.ReadPkg() //mes就是 RegisterResMes

	if err != nil {
		fmt.Println("readpkg(conn)err = ", err)
		return
	}

	//将mes的data部分反序列化成registerresmes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，重新登陆")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
