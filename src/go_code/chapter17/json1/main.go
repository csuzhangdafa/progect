package main
import (
	"fmt"
	"encoding/json"
)


type Monster struct{
    Name string   
	Age int 
	Birthdy string
	Sal float64
	Skill string
}


//将字符串反序列化为结构体
func unmarshalstruct(){

	//在项目开发中通过一些渠道获得
	str := "{\"Name\":\"牛魔王\",\"Age\":34,\"Birthdy\":\"2011/11\",\"Sal\":8000.3,\"Skill\":\"牛魔全\"}"//转义
    
	//定义一个monster的实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
    if err != nil{
		fmt.Printf("%v",err)
	}
	fmt.Printf("反序列化后=%v",monster)
}


//反序列化为序列化的逆操作unmarshal
func main(){
	unmarshalstruct()
}