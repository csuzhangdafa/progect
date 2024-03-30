/*基本介绍：闭包就是一个函数和与其相关的引用环境组合的一个整体(实体)
闭包让你可以在一个内层函数中访问到其外层函数的作用域。
可简单理解为：有权访问另一个函数作用域内变量的函数都是闭包。也可以说函数内部引用函数外的变量就是闭包
函数内部有一个变量，匿名函数是方法，他们都是一个整体，都在函数里面
*/

package main
import (
	"fmt"
	"strings"
)
//闭包
//累加器
func addupper() func (int)int{ //此函数的返回值为一个函数func，此函数的返回值为int，遵循函数性质8
	var n int =10
	var str = "hello"
	return func (x int)int{
		str += "a"
		n = n+x
		fmt.Println("str=",str)
		return n
	}
}

func makesuffix(suffix string, name string)string{


		if !strings.HasSuffix(name, suffix) {
		    return name + suffix
		}
		return name 

}
func main(){
	f :=addupper()         //调用累加器，返回值为一个函数。这时就可以调用
	fmt.Println(f(1))      //11
	fmt.Println(f(2))      //13形成闭包后，函数内部的变量已经改变
	fmt.Println(f(3))      //16

//1.addupper是一个函数，返回的数据类型是fun(int)int
//2.闭包的说明：返回的是一个匿名函数，但是这个匿名函数引用到了函数外的n，这两个就形成一个整体，构成了闭包。
//3.闭包是类，函数是操作，n是字段
//4.当我们反复调用函数f时，因为n是初始化一次，因此每调用一次就进行累计
//5.搞清楚闭包的关键，就是要分析出返回的函数使用哪些变量，因此函数和它用到的变量形成闭包
//会，因为使用闭包会包含其他函数的作用域，会比其他函数占据更多的内存空间，不会在调用结束之后被垃圾回收机制回收，
多度使用闭包会过度占用内存，造成内存泄漏。
	fmt.Println("文件名处理后=",makesuffix(".jpg","hello"))
}