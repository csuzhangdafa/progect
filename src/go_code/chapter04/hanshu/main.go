/*
func 函数名(参数列表)(返回值列表){
	执行语句
	return 返回值列表
}
1  函数的形参列表可以是多个，返回值列表也可以是多个。
2  形参列表和返回值列表的数据类型可以是值类型和引用类型。
3  函数的命名遵循标识符命名规范，首字母不能是数字，
首字母大写该函数可以被本包文件和其它包文件使用，类似public, 首字母小写，
只能被本包文件使用，其它包文件不能使用，类似privat
4  函数中的变量是局部的，函数外不生效
5  基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值。
6. 如果希望函数内的变量能修改函数外的变量(指的是默认以值传递的方式的数据类型)，可以传入变量的地址&，
函数内以指针的方式操作变量。从效果上看类似引用 。
7  Go函数不支持函数重载
8  在Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
9  函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
10  为了简化数据类型定义，Go支持自定义数据类型 基本语法：type 自定义数据类型名 数据类型
// 理解: 相当于一个别名



1  如果一个文件同时包含全局变量定义， init 函数和 main 函数，则执行的流程全局变量定义 - >init函数 - >main 函数
2  init函数最主要的作用，就是完成一些初始化的工作
package main
import(
	"fmt"
)
var age = test()

func test() int {
	fmt.Println("test()") //1
	return 90
}

func init(){
	fmt.Println("init()") //2
}

func main(){
	fmt.Println("main()...age=",age) //3
}

Go支持匿名函数，匿名函数就是没有名字的函数，如果我们某个函数只是希望使用一次，可以考 虑使用匿名函数，
匿名函数也可以实现多次调用。

匿名函数使用方式 1
在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次

package main
import (
	"fmt"
)
func main() {
	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1=", res1)
}

匿名函数使用方式 2
将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函

package main
import (
	"fmt"
)

func main() {

	//将匿名函数func (n1 int, n2 int) int赋给 a变量
	//则a 的数据类型就是函数类型 ，此时,我们可以通过a完成调用
	a := func (n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 30)
	fmt.Println("res2=", res2)
	res3 := a(90, 30)
	fmt.Println("res3=", res3)
}

全局匿名函数
如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以在程序有效。

defer函数
package main
import(
	"fmt"
)

func sum(n1 int, n2 int)int{
	当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
	defer fmt.Println("ok1 n1=",n1)
	defer fmt.Println("OK2 N2=",n2)

	res := n1+n2
	fmt.Println("ok3 res=",res)
	return res
}

func main(){
	res:=sum(10,20)
	fmt.Println("res=",res)
}
*/

package main

func main() {
}
