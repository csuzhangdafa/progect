package main
import (
	"fmt"
)


//类型断言最佳实践用phone结构体增加一个特有方法call()当usb接口接到phone时还需要调用call方法


type Usb interface{  //名字为Usb的接口
	Start()          //接口内部的方法
	Stop()
}


type Phone struct{
    name string
}


type Camera struct{
    name string
}


type Computer struct{

}


//让phone实现Usb的方法
func (p Phone) Start(){
	fmt.Println("手机开始工作")
}


func (p Phone) Stop(){
	fmt.Println("手机停止工作")
}


func (p Phone) Call(){
	fmt.Println("手机开始打电话")
}


//让相机实现接口的方法
func (c Camera) Start(){
	fmt.Println("相机开始工作")
}


func (c Camera) Stop(){
	fmt.Println("相机停止工作")
}


func (computer Computer) working(usb Usb){
	usb.Start()
	//如果usb是指向phone的结构体变量，则还需要调用call方法
	//类型断言
	if phone , ok := usb.(Phone); ok{
		phone.Call()
	}
	usb.Stop()
}


//编写一个函数判断传入的参数是什莫类型
func Typejudge(items... interface{}){   //输入参数的值不确定可以这样写
    for index, x := range items{


		switch x.(type){
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n",index+1, x)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n",index+1, x)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是%v\n",index+1, x)
		case int,int32,int64:
			fmt.Printf("第%v个参数是 整数 类型，值是%v\n",index+1, x)
		case string:
			fmt.Printf("第%v个参数是 string类型，值是%v\n",index+1, x)
		case Student:
			fmt.Printf("第%v个参数是 Student类型，值是%v\n",index+1, x)
		case *Student:
			fmt.Printf("第%v个参数是 *Student类型，值是%v\n",index+1, x)
		default :
		fmt.Printf("第%v个参数是类型不确定，值是%v\n",index+1, x)
		}


	}
}


type Student struct{

}


func main(){

	var usbarr [3]Usb
	
	usbarr[0] = Phone{"vivo"}
	usbarr[1] = Phone{"小米"}
	usbarr[2] = Camera{"尼康"}


	var computer Computer
	for _, v := range usbarr{    //index为下标，v为元素
        computer.working(v)
	}

	//fmt.Println(usbarr)

	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var n4 int64 = 800
	var n5 string = "tom"
	address := "北京"
	n6 := 300
	stu1 := Student{}
	stu2 := &Student{}
	Typejudge(n1,n2,n3,n4,n5,address,n6,stu1,stu2)
}