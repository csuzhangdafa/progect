package main
import (
	"fmt"
	"io/ioutil" //使用os包也可
)


//使用一次性读入文件，适用于文件不太大的情况
func main(){

    //使用iouti.ReadFile一次性将文件读取到位
	file := "d:/text.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil{
		fmt.Printf("read file err=%v",err)
	} 
	fmt.Printf("%v",string(content))

	//没有打开文件和关闭文件是因为将封装到了函数内部

}