package main
import (
	"fmt"
)


/*类型断言由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
var a interface{}
var point Point = Point{1,2}
a = point
var b Point
b = a //error
b = a.(Point) //表示判断a是否指向Point类型的变量，如果是就转成Point类型并给b赋值，否则就报错*/


func main(){ 
    var x interface{}
	var b float32 = 1.1
	x = b 
	//x转为float32使用类型断言
	y , ok:= x.(float32)
	if ok{         //if y,ok :=x.(float32); ok{}
		fmt.Println("转换成功")
		fmt.Printf("y的类型为 %T 值为 %v",y,y)
	}else{
		fmt.Println("转换失败")
	}


	//在断言时带一个检查机制做一个判断语句防止断言不成功时程序终止

}