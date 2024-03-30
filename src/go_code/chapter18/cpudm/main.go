package main
import (
	"fmt"
	"runtime"
)


func main(){
	cpuNum := runtime.NumCPU()   //检测当前有多少个CPU
	fmt.Println("当前cpu个数为", cpuNum)


	//也可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)//设置的数量预留一个
	fmt.Println("OK")
}