package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)


//读取文件内容在终端显示,使用带缓冲区的方式,可以读取较大的文件
func main(){
	
	//打开一个文件
	file , err := os.Open("d:/text.txt")
	if err != nil{
		fmt.Println("open file err = ",err)
	}

    //函数退出时及时关闭file
	defer file.Close()  //及时关闭file句柄，否则会有内存泄漏
	//创建一个reader，带缓冲
	//默认缓冲区为4096
    reader := bufio.NewReader(file)
	//循环读取文件内容
	for{
		str, err := reader.ReadString('\n') //读到换行就结束一次 , str是读取的信息，err是错误反馈
		if err == io.EOF{   //表示文件的末尾
            break
		}

		//输出内容
		fmt.Print(str)
	}

	fmt.Println("文件读取结束")
}