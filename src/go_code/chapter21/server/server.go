/*
服务端的处理流程
1 监听端口 8888
2 接收客户端的tcp链接，建立客户端和服务器端的链接
3 创建 goroutine ，处理该链接的请求(通常客户端会通过链接发送请求包)

客户端的处理流程
1 建立与服务端的链接
2 发送请求数据[终端]，接收服务器端返回的结果数据
3 关闭链接

defer 入栈
*/

package main

import(
	"fmt"
	"net"
	_"io"
)

//首先要拿到接口
func process(conn net.Conn){

	defer conn.Close()//关闭接口
	for{
		//创建一个新的切片来读取客户端传来的信息	
		buf := make([]byte,1024)

		//conn.Read(buf)
		//1. 等待客户端通过conn发送信息
		//2. 如果客户端没有wrtie[发送]，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("客户端退出 err=",err)

			//此处因为客户端的退出服务端出现了一个错误，在不断闪屏，应该及时推出协程
			return
		}

		//显示客户端发送的内容到终端，n代表实际读到的东西,循环等待
		fmt.Print(string(buf[:n]))
	}
}

func main(){

	/*服务端功能
      编写一个服务器端程序，在 8888 端口监听
      可以和多个客户端创建链接
      链接成功后，客户端可以发送数据，服务器端接受数据，并显示在终端上
      先使用telnet 来测试，然后编写客户端程序来测试*/
	fmt.Println("服务器开始监听")
	//net.Listen("tcp", "0.0.0.0:8888")
	//1. tcp 表示使用网络协议是tcp
	//2. 0.0.0.0:8888 表示在本地监听 8888端口
	listen,err:=net.Listen("tcp","0.0.0.0:8888")
	/*
	func Listen(net, laddr string) (Listener, error)
	返回在一个本地网络地址laddr上监听的Listener。网络类型参数net必须是面向流的网络

	type Listener interface {
        // Addr返回该接口的网络地址
        Addr() Addr
        // Accept等待并返回下一个连接到该接口的连接
        Accept() (c Conn, err error)
        // Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
        Close() error
    }


	type Addr interface {
        Network() string // 网络名
        String() string  // 字符串格式的地址
    }


	type Conn interface {
        // Read从连接中读取数据
        // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
        Read(b []byte) (n int, err error)
        // Write从连接中写入数据
        // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
        Write(b []byte) (n int, err error)
        // Close方法关闭该连接
        // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
        Close() error
        // 返回本地网络地址
        LocalAddr() Addr
        // 返回远端网络地址
        RemoteAddr() Addr
        // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
        // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
        // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
        // 参数t为零值表示不设置期限
        SetDeadline(t time.Time) error
        // 设定该连接的读操作deadline，参数t为零值表示不设置期限
        SetReadDeadline(t time.Time) error
        // 设定该连接的写操作deadline，参数t为零值表示不设置期限
        // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
        SetWriteDeadline(t time.Time) error
    }
	*/

	if err !=nil{
		fmt.Println("listen err=",err)
		return
	}
	defer listen.Close()  //延时关闭listen

	//端口已经监听，等待客户端的连接
	for {
		fmt.Println("等待客户端连接..")
		conn,err:=listen.Accept()
		if err!=nil{

			//如果err不为空则证明没连接
			fmt.Println("Accept() err=",err)
		}else{

			//err为空则证明连接
			fmt.Printf("Accept() suc con = %v 客户端ip=%v 网络名=%v\n ",conn,conn.RemoteAddr().	String(),conn.RemoteAddr().Network())
		}

		//此处应该起协程为客户端服务
		go process(conn)
	}
}