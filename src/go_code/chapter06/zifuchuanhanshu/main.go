package main
import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	//字符串长度
	str := "hello北"
	fmt.Println("str len =",len(str))


    //字符串中全部转为字符串类型
	str2 := "hello北京"
	r :=[]rune(str2)
	for i :=0; i<len(r);i++{
		fmt.Printf("字符=%c\n",r[i])
	}

    
	//字符串转整数
	n , err :=strconv.Atoi("123")
	if err !=nil{
		fmt.Println("模式转换错误")
	}else{
		fmt.Println("转换的结果是\n",n)
	}


	//整数转字符串
	str3 :=strconv.Itoa(12345)
	fmt.Printf("str3=%v,str3=%T\n",str3,str3)


    //字符串转切片
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)


	//切片转字符串
	str4 := string([]byte{97,98,99})
	fmt.Printf("str4=%v\n",str4)


	//十进制转
	str5 := strconv.FormatInt(123,2)
	fmt.Println("123对应的二进制是=%v\n",str5)
	str6 := strconv.FormatInt(123,16)
	fmt.Println("123对应的十六进制是=%v\n",str6)


	//查找子串是否在指定的字符串中
    b :=strings.Contains("seafood","marry")
	fmt.Printf("b=%v\n",b)
	

	//统计一个字符串中有几个指定的字串
	num :=strings.Count("cehese","c")
	fmt.Printf("num=%v\n",num)


    //不区分字母大小写的比较
	a := strings.EqualFold("abc","ABC")
	fmt.Printf("a=%v\n",a)


	//区分大小写的比较
	fmt.Println("结果\n","abc"=="ABC")


	//返回字串在字符串中出现的的一个index值，如果没有有，返回-1
	index := strings.Index("asdfghjdgf","jdg")
	fmt.Printf("index=%v\n",index)


	//返回字符串最后一个index
	lastindex := strings.LastIndex("ghjdgfsddfjdg","jdg")
	fmt.Printf("lastindex=%v\n",lastindex)


    //将指定字串替换成另一个字串若n=-1则表示全部替换n表示替换几个字符串
	str7 := strings.Replace("go go hello","go","北京",2)
	fmt.Println("str7=\n",str7)



}