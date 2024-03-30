package main
import (
	"fmt"
	"io/ioutil"
)



func main(){

	//将d：adc导入到e：kkk
	//1.读取abc到内存 2.读取到的内容写入

	file1path := "d:/abc.txt"
	file2path := "e:/kkk.txt"
	data, err := ioutil.ReadFile(file1path)
	if err != nil{
		fmt.Println("read file err")
		return 
	}


	err = ioutil.WriteFile(file2path, data, 0666)
	if err != nil{
		fmt.Println("write file error")
	}

}