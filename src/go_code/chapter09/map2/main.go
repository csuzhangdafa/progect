package main
import (
	"fmt"
)


//使用一个函数，用map类型如果用户名存在将其密码修改为888888，不存在就添加信息

func modifyuser(users map[string]map[string]string,name string){
    if users[name] !=nil{
		users[name]["pws"]="888888"
	}else{
		users[name] = make(map[string]string, 2)
		users[name]["pwd"]="888888"
		users[name]["nickname"]="昵称" + name
	}
}

func main(){
	users := make(map[string]map[string]string,10)
	users["smith"]=make(map[string]string,2)
	users["smith"]["pwd"]="999999"
	users["smith"]["nickname"]="小花猫"
	modifyuser(users,"tom")
	modifyuser(users,"mary")
	modifyuser(users,"smith")
	fmt.Println(users)
}