/*
func 函数名(参数列表)(返回值列表){
	执行语句
	return 返回值列表
}函数定义

func(recevier type)方法名（参数列表） (返回值列表){
    方法体
    return 返回值
}  方法定义

方法与函数的区别：
函数的调用方式: 函数名(实参列表)
方法的调用方式: 变量.方法名(实参列表)

*/

package main
import (
	"fmt"
)


//golang中的方法是作用在指定数据类型上的（即：和指定的数据类型绑定因此自定义数据类型都可以有方法不仅仅是struct


type person struct{
	name string
}


//给A类型绑定一份方法 p相当于persion的变量名
func (p person) test(){
	p.name = "jack"
    fmt.Println("test() ",p.name)
}


func main (){
    var  p person
	p.name = "tom"
	p.test()   //调用方法
    fmt.Println("main() p.name= ",p.name)   //传参并不影响主函数的值
}


//1.test方法和person绑定
//2.test方法只能通过person变量类型来调用，而不能直接调用，也不能使用其他类型的变量来调用