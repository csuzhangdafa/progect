package main
import (
	"fmt"
)


//方法调用和传参机制
//方法调用时会将调用方法的变量，当作实参也传递给方法
//如果变量是值类型，则进行拷贝，如果是引用类型则进行地址拷贝


//方法快速入门案例1.给person结构体添加一个speak方法，输出，xxx是一个好人s
type person struct{
	name string
}


func (p person) speak(){  //方法和函数相似内部可以接受函数 speak(n int, , , )
	fmt.Println(p.name,"is a good man") 
}


//对其它类型使用方法
type integer int
func (i *integer)print(){
	*i=*i+1
	fmt.Println("i=",*i)
}


//如果一个类型实现了string()这个方法，那么fmt.Println默认会调用这个变量的string()进行输出
type student struct{
	name string
	age int
}


//给student实现string方法
func (stu *student) string() string{ //第一个string为方法名，第二个string为返回值
    str := fmt.Sprintf("name=[%v] age=[%v]",stu.name,stu.age)
	return str
}


func main(){
    var v person
	v.name = "tom"
	v.speak()
	var q integer = 8
	q.print()

	stu := student{       //结构体赋值与map相似
		name : "tom",
		age : 20,
	}
	fmt.Println(&stu)
}




 //对于普通函数，接受者为值类型，不能将指针类型的数据类型传递，反之亦然
 //对于方法，接受者为值类型可以直接用指针类型的变量调用方法，反过来也同样可以