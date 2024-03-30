/*
1.在默认情况下，当发生错误后(panic),程序就会退出（崩溃.）
2.如果我们希望：当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。
还可以在捕获到错误后，给管理员一个提示(邮件,短信。。。）


1.Go语言追求简洁优雅，所以，Go语言不支持传统的 trycatchfinally 这种处理。
2.Go中引入的处理方式为： defer , panic , recover
3.这几个异常的使用场景可以这么简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，
然后正常处理


Go程序中，也支持自定义错误， 使用errors.New 和 panic 内置函数。
1.errors.New(“错误说明”), 会返回一个error类型的值，表示一个错误
2.panic内置函数 ,接收一个interface{}类型的值（也就是任何值了）作为参数。
可以接收error类型的变量，输出错误信息，并退出程序.


*/


package main
import (
	"fmt"
	"errors"
)



//错误处理defer panic recover
func test(){
	//使用defer+recover来捕获和处理异常  defer见04/hanshu
	defer func(){
		err := recover()//内置函数可以捕获异常
		if err !=nil{//说明捕获到错误
            fmt.Println("err=",err)
			//将错误信息发送给管理员
		}
	}()
	num1 := 10
	num2 := 0
	res :=num1 / num2
	fmt.Println("res=",res)
}


//自定义错误
//函数去读取以配置文件init.conf的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string)(err error){
    if name =="config.ini"{
		//读取
		return nil      //文件名正确，没有错误返回一个空值
	}else{
		return errors.New("读取文件错误")//返回一个自定义错误
	}
}

func test02(){

	err :=readConf("config.ini")
	if err !=nil{
		//读取文件错误抛出panic终止程序
		panic(err) //内置函数直接用终止程序
	}

	fmt.Println("test02()继续执行")
}
func main(){
	/*test()
	fmt.Println("main()下面的代码...")*/

	test()
	test02()
	fmt.Println("main()下面的代码...")
}