package main
import (
    "fmt"
	"io"
	"bufio"
	"os"
)


//自己编写一个函数接收两个文件的路径
func Copyfile(desfilename string, srcfilename string)(written int64, err error){
    
	srcfile, err := os.Open(srcfilename)
	if err != nil{
		fmt.Printf("%v\n",err)
	}
    defer srcfile.Close()


	//通过srcfile句柄。获取reader
    reader := bufio.NewReader(srcfile)


	//打开desfilename，如果不存在就创建
	desfile, err := os.OpenFile(desfilename, os.O_WRONLY | os.O_CREATE, 0666)
	if err!=nil{
		fmt.Printf("%v\n",err)
		return
	}
	defer desfile.Close()


    //通过desfile句柄。获取writer
    writer := bufio.NewWriter(desfile)


	return io.Copy(writer, reader)

}


func main(){
	
	//拷贝文件

	//调用函数
	srcfile := "d:/flower.jpg"
	desfile := "e:/abc.jpg"
	_, err := Copyfile(desfile, srcfile)
	if err==nil{
		fmt.Printf("拷贝完成\n")
	}else{
		fmt.Printf("%v",err)
	}
}