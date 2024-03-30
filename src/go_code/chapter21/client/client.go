package main
import(
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main(){
	/*
	Dial函数和服务端建立连接：

    conn, err := net.Dial("tcp", "google.com:80")//服务器IP加端口
    if err != nil {
	    // handle error
    }
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    status, err := bufio.NewReader(conn).ReadString('\n')//创建一个reader可以接收文件的输入
	*/
    conn,err:=net.Dial("tcp","192.168.117.123:8888")
	if err != nil{
		fmt.Println("client dial err=",err)
		return
	}
	fmt.Println("conn 成功= ",conn)

	/*功能一：客户端可以发送单行数据，然后就退出
	NewReader创建一个具有默认大小缓冲、从r读取的*Reader
	os.Stdin 作为输入源。os.Stdin 表示标准输入流
	*/
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入

	/*
	func (*Reader) ReadString
    func (b *Reader) ReadString(delim byte) (line string, err error)
    ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。
    如果ReadString方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。
    当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。
	从标准输入中读取一行文本，然后根据可能的错误输出错误信息 

	
	从终端读取一行用户输入，并准备发给服务器
	*/
    for{


		line,err := reader.ReadString('\n')
		if err!= nil{
			fmt.Println("readstring err=",err)
		}


		//如果用户输入的是exit就退出
		line = strings.Trim(line,"\r\n")
		if line == "exit"{
			fmt.Println("客户端退出")
			break
		}

		//将line发送给服务器
		n,err:=conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.write err=",err,n)
		}
	}
}