package main
import (
	"fmt"
)


//多态是通过接口实现的，可以按照统一的接口来调用不同的实现，这时接口变量就呈现出不同的形态
//多态数组



//定义一个接口
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


//让相机实现接口的方法
func (c Camera) Start(){
	fmt.Println("相机开始工作")
}


func (c Camera) Stop(){
	fmt.Println("相机停止工作")
}


/*让计算机工作的方法
func (c Computer)working(usb Usb){ //计算机绑定一个可以接受Usb接口的方法
	usb.Start()
	usb.Stop()
}*/


func main(){
    /*computer := Computer{}
    Phone := Phone{}
	Camera := Camera{}

	computer.working(Phone)
	computer.working(Camera)*/

	//定义一个usb接口可以存放，phone 和 camera 的结构体变量
	//体现多态数组


	var usbarr [3]Usb
	
	usbarr[0] = Phone{"vivo"}
	usbarr[1] = Phone{"小米"}
	usbarr[2] = Camera{"尼康"}
	fmt.Println(usbarr)

}
