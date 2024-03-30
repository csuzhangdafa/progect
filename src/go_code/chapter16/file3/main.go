package main
import (
	"fmt"
	"bufio"
	"os"
)



func main(){

	//创建一个新文件，并在文件中写入内容
	//1.打开文件,不存在的文件直接创建
	filePath := "d:/abc.txt"
    file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)
    if err!=nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}


	//及时关闭，防止泄露
	defer file.Close()

	//打开成功，写入内容
	str := "你好\n"
	//写入时使用带缓存的writter
	writer := bufio.NewWriter(file)
	for i:= 0; i<10; i++{
		writer.WriteString(str)
	}

	//因为Writer是带缓存的，因此在调用方法时是先写入缓存的,使用Flush方法写入文件中
	writer.Flush()
}