package main
import (
	"fmt"
	"bufio"
	"os"
)



func main(){

	//打开一个文件将原来的内容覆盖成10句 你好
	//在原来的内容追加ABC
	filePath := "d:/abc.txt"
    //覆盖   file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
    //追加   file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
    file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
    if err!=nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}


	//及时关闭，防止泄露
	defer file.Close()

	//打开成功，写入内容
	str := "ABC\n"
	//写入时使用带缓存的writter
	writer := bufio.NewWriter(file)
	for i:= 0; i<5; i++{
		writer.WriteString(str)
	}

	//因为Writer是带缓存的，因此在调用方法时是先写入缓存的,使用Flush方法写入文件中
	writer.Flush()
}