/*工厂模式
如果定义的结构体首字母是大写可以直接跨包引用，如果是小写就要用到工厂模式

*/
package main

import (
	"fmt"
	"go_code/chapter12/severce"
)

func main(){
	var stu = severce.NewStudent("tom",98.8) 
	fmt.Println(*stu)
	fmt.Println("name=",stu.Name,"score=",stu.Getscore())
}

//包注意细节。 1.在go_project下go mod init 包名
//2.引用  包名/文件名
//如果结构体首字母为大写则可以直接引用
//使用工厂模式实现跨包创建结构体变量首字母小写