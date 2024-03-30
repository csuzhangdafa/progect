package main
import (
	"fmt"
	"time"
	"strconv"
)


func test03(){
	str :=""
	for i :=0;i<100000;i++{
		str +="hello"+strconv.Itoa(i)
	}
}
func main(){
    
	//获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n",now,now)
    

	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v\n",int(now.Month()))
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())

	//格式化日期时间
	fmt.Printf(now.Format("2006/01/02 15:04:05"))


	//Unix UnixNano时间戳方法获取当前时间
	fmt.Printf("unix时间戳=%v unixnano=%v\n",now.Unix(),now.UnixNano())
     


	//统计函数执行的时间
	start := time.Now().Unix()
    test03()
	end := time.Now().Unix()
	fmt.Printf("执行耗费的时间为%v秒\n",end-start)

}