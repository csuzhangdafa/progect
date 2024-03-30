/*interface类型可以定义一组方法，但是这些都不需要实现，并且interface不能包含任何变量。
到某个自定义类型需要使用时（比如结构体），根据具体情况把这些方法写出来
golang内的接口实现是基于方法，并不是接口名称


type 接口名 interface{

	method1(参数列表)返回值列表
	method(2参数列表)返回值列表
	.....

}


实现接口的方法


func (t 自定义类型)method1(参数列表)返回值列表{
	//方法实现
}


func (t 自定义类型)method2(参数列表)返回值列表{
	//方法实现
}


使用细节
1.接口本身不能创建实例，但是可以指定一个实现该接口的自定义类型的变量
2.接口中所有的方法都实现才能说实现了这个接口
3.一个自定义类型可以实现多个接口
4.接口可以继承，要实现接口就要实现这个结构继承的所有接口
5.interface默认是一个指针（引用类型）如果没有初始化就输出nil
6.如果为空接口，所有的类型都实现了空接口


*/

package main
import (
	"fmt"
)


//多态的特性主要是依靠接口来实现，所以在学习多态之前先学习接口


//快速入门
//定义一个接口
type Usb interface{  //名字为Usb的接口
	Start()          //接口内部的方法
	Stop()
}


type Phone struct{

}


type Camera struct{

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


//让计算机工作的方法
func (c Computer)working(usb Usb){ //计算机绑定一个可以接受Usb接口的方法
	usb.Start()
	usb.Stop()
}


func main(){
    computer := Computer{}
    Phone := Phone{}
	Camera := Camera{}

	computer.working(Phone)
	computer.working(Camera)
}









