package main
import (
	"fmt"
	"os"
)



func main(){
    //打开一个文件
	file , err := os.Open("d:/text.txt")
	if err != nil{
		fmt.Println("open file err = ",err)
	}

	fmt.Printf("file=%v",file)

	err = file.Close()
	if err != nil{
		fmt.Println("close file err = ",err)
	}
}