/*channle本质就是一个数据结构-队列
数据是先进先出【FIFO:firstinfirstout】
线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
channel有类型的，一个string的channel只能存放string类型数据。


定义声明channel

var 变量名 chan 数据类型

举例：
var intChan chan int(intChan用于存放int数据)
var mapChan chan map[int]string(mapChan用于存放map[int]string类型)
var perChan chan Person
var perChan2 chan *Person

channel是引用类型，存储的是地址
channel必须初始化才能写入数据, 即make后才能使用
管道是有类型的，intChan 只能写入 整数 int

使用内置函数close可以关闭channel, 当channel关闭后，就不能再向channel写数据了，但是仍然 可以从该channel读取数据
close(intChan) // close

channel支持for–range的方式进行遍历
在遍历时，如果channel没有关闭，则会出现deadlock的错误
在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历

你可以使用for range语句来遍历切片（slice）、数组（array）、字符串（string）、映射（map）和通道（channel）等数据结构
for index, value := range collection {
    // 在这里处理index和value
}

*/

package main

import (
	"fmt"
)

func main() {

	//定义，make以后使用
	var intchan chan int
	intchan = make(chan int, 3)
	fmt.Printf("intchan的值=%v\n", intchan)

	//写入数据
	intchan <- 10
	num := 211
	intchan <- num

	fmt.Printf("channel len= %v cap=%v \n", len(intchan), cap(intchan)) //长度2，容量3

	//读取数据，取数据的时候也可以不接收
	var num2 int
	num2 = <-intchan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intchan), cap(intchan))

	//创建一个mapChan，最多可以存放10个map[string]string的key-val，演示写入和读取
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	//map使用前先make
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "天津"

	m2 := make(map[string]string, 20)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"

	mapChan <- m1
	mapChan <- m2

	m11 := <-mapChan
	m22 := <-mapChan

	fmt.Println(m11, m22)

}
