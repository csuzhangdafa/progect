/*
在Golang中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性。

*/
package main
import (
	"fmt"
)


type A struct{
	Name string
	age int
}


func (a *A)Sayok(){
	fmt.Println("A Sayok",a.Name)
}


func (a *A)hello(){
	fmt.Println("A hello",a.Name)
}


type B struct{
	A
}


//继承深入讨论
//如果嵌入两个结构体，使用的时候就需要指定是哪一个
//如果一个结构体内嵌入一个有名结构体，这种模式称为组合或者多重继承，如果是组合关系，那么在访问结构体时，必须带上结构体的名字
//基本数据类型也可以嵌入到结构体中
//为了代码的简洁性，建议不要使用多重继承


func main(){
	var b B 
	b.Name = "tom"
	b.age = 19
	b.Sayok()
	b.hello()
}