/*
Go 语言中有两种主要的数据类型：值类型和引用类型。值类型包括基本数据类型（如整数、浮点数、布尔值等）以及结构体（struct）。
引用类型则包括切片（slice）、映射（map）、通道（channel）、接口（interface）和函数。

引用类型都是通过指针来实现的，即变量存储的是数据的地址，而不是实际的数据。

切片的长度是可以变化的，因此切片是一个可以动态变化数组

定义:
var 切片名 []类型
var a []int

1.通过make来创建切片,此方法创建要求cap容量>=len长度
var 切片名 []type = make([]type,len,[cap])

2.res := [][]int{{}, {}}
这将创建一个包含两个空整数切片的二维切片。

3.res := [][]int{}
res = append(res, []int{1, 2, 3})
这将创建一个空的二维整数切片，并向其中添加了一个包含 {1, 2, 3} 的整数切片。

slice从底层来说就是一个数据结构
type slice struct{
	ptr *int        //指向底层数组的首地址
	len int         //切片长度
	cap int         //容量
}
*/

package main

import (
	"fmt"
)

// 切片使用案例
func fbn(n int) []uint64 {
	//申明一个切片大小为n
	fbnslice := make([]uint64, n)
	fbnslice[0] = 1
	fbnslice[1] = 1
	for i := 2; i < n; i++ {
		fbnslice[i] = fbnslice[i-1] + fbnslice[i-2]
	}
	return fbnslice
}

func main() {
	//定义 var 变量名 []类型  比如var a []int
	var intarr [5]int = [...]int{1, 22, 33, 44, 66}

	//1.定义一个切片slice引用到intarr数组第一个到第三个
	slice := intarr[1:3]
	fmt.Println("intarr=\n", intarr)
	fmt.Println("slice的元素是\n", slice)
	fmt.Println("slice len\n", len(slice))

	//2.内置函数make创建切片,对于切片必须make以后才可以使用,此方式创建的切片在底层进行维护，程序员看不到,只能通过slice去访问各个元素
	var slice01 []float64 = make([]float64, 5, 10) //长度，容量
	slice01[1] = 10
	slice01[3] = 34
	fmt.Println("\n", slice01)
	fmt.Println("slice01长度\n", len(slice01))
	fmt.Println("slice01容量\n", cap(slice01))

	//3.定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var slice02 []string = []string{"tom", "jack", "marry"}
	fmt.Println("\n", slice02)
	fmt.Println("长度", len(slice02))
	fmt.Println("容量", cap(slice02))

	//切片的遍历
	//1.常规的for循环遍历
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice03 := arr[:]
	for i := 0; i < len(slice03); i++ {
		fmt.Printf("slice03[%v]=%v\n", i, slice03[i])
	}

	//for range 方式遍历切片
	for i, v := range slice03 {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	//append对切片动态增加,是把追加以后的切片形成一个新的切片重新赋值，原先的切片值并没有改变
	//切片append操作的本质就是对数组扩容
	//go底层会创建一下新的数组newArr(安装扩容后大小)
	//将slice原来包含的元素拷贝到新的数组newArr
	//slice 重新引用到newArr
	//注意newArr是在底层来维护的，程序员不可见.
	slice04 := []int{100, 200, 300}
	slice04 = append(slice04, 400, 500, 600) //追加的切片名，追加的切片元素 100,200,300,400,500,600
	slice04 = append(slice04, slice...)      //可以追加本身,或者别的切片
	fmt.Println("slice04=", slice04)

	//切片的拷贝copy(par1,par2)//参数的数据类型为切片
	var slice05 []int = []int{1, 2, 3, 4, 5}
	var slice06 = make([]int, 10)
	copy(slice06, slice05)
	fmt.Println("slice05=\n", slice05)
	fmt.Println("slice06=\n", slice06)

	//slice为引用类型传递变量值都会改变

	//string底层是一个byte数组，可以进行切片处理
	str := "hello@atguigu"
	slice07 := str[6:]
	fmt.Printf("str=%v,slice07=%v\n", str, slice07)

	//字符串不可改，string是不可变的，也就说不能通过 str[ 0 ]=‘z’ 方式来修改字符串.如果需要改先转为切片改完以后再转为字符串
	slice08 := []rune(str) //arr1 := []byte(str)也可以转换成字节切片
	slice08[0] = '北'       //转为切片后不能处理中文,降string转为[]rune即可
	str = string(slice08)
	fmt.Println("str=\n", str)

	//切片案例斐波那契
	fbnslice := fbn(20)
	fmt.Println("fbnslice=", fbnslice)
}
