package main
import (
	"fmt"
)
type myfuntype func(int,int) int
func getSun(n1 int, n2 int) int {
	return n1+n2
}

func myFun(funvar func(int,int) int,num1 int,num2 int) int {
	return funvar(num1,num2)
}
func myFun2(funvar myfuntype,num1 int,num2 int) int {
	return funvar(num1,num2)
}
func getsumandgetsub(n1 int,n2 int)(sum int,sub int){
    sum = n1 + n2
	sub = n1 - n2
	return 
}

func main(){
    a:= getSun
	fmt.Printf("a的类型是%T,getSun类型是%T\n", a , getSun)

	res := a(10,40)
	fmt.Println("res=",res)

	res3 := myFun(getSun,500,600)
	fmt.Println("res3=",res3)
	type myint int
	var num myint = 40
	var num2 int
	num2 = int(num)
	fmt.Println("num=,num2=",num,num2)

    a1,b1 := getsumandgetsub(1,2)
	fmt.Println("a1=%v,b1=%v\n",a1,b1)

}
