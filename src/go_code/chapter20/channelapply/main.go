package main
import(
	"fmt"
	"time"
)

/*
1.开启一个wirtedata协程，向管道中写入50个整数
2.开启一个readdata协程，向管道中读取写入的数据
两个协程为同一个管道
主线程需要等待write和read协程都完成工作时再推出(管道)
*/

//开日write协程去往管道内写东西
func writedata(intchan chan int){
	for i:= 0;i<50;i++{
		/*注意此时要先打印数据在写入，但在循环内部却直接使用 
		fmt.Println 打印了一些内容，而没有先完全写入整数到通道中。这导致了协程的写入和读取的速度不同步，
		可能会造成读取协程读取到的数据比预期少。*/
		fmt.Println("readdata",i)
		intchan<-i
		time.Sleep(time.Second)
	}
	close(intchan) //写完关闭
}

//开启read协程往管道内读取东西
/*v: 这是一个变量，用于存储从通道 intchan 接收到的数据。
ok: 这也是一个变量，用于存储一个布尔值，表示通道是否已关闭。通道关闭后，
再从通道中接收数据会得到零值，并且 ok 会被设置为 false，否则为 true
*/
func readdata(intchan chan int,exitchan chan bool){
	for{
		v, ok :=<-intchan
		if !ok{
			break
		}
		fmt.Printf("读到的数据=%v\n", v)
		time.Sleep(time.Second)
	}
	//readdata读取完数据后，任务完成
	exitchan<-true
	close(exitchan) 
}

func main(){

	//创建两个管道
	intchan := make(chan int,50)
	exitchan := make(chan bool, 1)
 
	go writedata(intchan)
	go readdata(intchan,exitchan)
	//此时如果不加代码约束，是秒闪，主线程结束了，协程也随之结束

	//time.Sleep(time.Second*10)//用休眠函数可以展现

	//用for去读代码退出的管道，防止主线程结束
	for{
		_, ok:=<-exitchan
		if !ok{
			break
		}
	}
}