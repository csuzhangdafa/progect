package main
import (
	"fmt"
	"sort"
	"math/rand"
)


//接口最佳实践 实现对hero结构体切片的排序 sort.Sort(data interface)
//使用系统提供的方法,接受的参数为一个接口,要实现接口的方法才可以调用这个方法


type Hero struct{
	Name string
	Age int 
}


//声明一个Hero结构体类型切片
type HeroSlice []Hero


//实现Interface接口,实现其中的三个方法
func (hs HeroSlice)Len() int{
	return len(hs)
}


//less方法就是决定你使用什么标准进行排序
//按Hero年龄从小到大排序
func (hs HeroSlice)Less(i,j int) bool{
	return hs[i].Age < hs[j].Age
}


func (hs HeroSlice)Swap(i,j int){
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}


func main(){
    var intSlice = []int{0,-1,10,7,90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)


	var heroes HeroSlice
	for i :=0 ; i<10; i++{
        hero := Hero{
			Name : fmt.Sprintf("英雄 %d",rand.Intn(100)),
            Age : rand.Intn(100),
		}

		
		//将hero append 到 heroes切片
		heroes = append(heroes,hero)
	}


	//看看排序前的顺序
	for _,v := range heroes{
		fmt.Println(v)
	}


	//调用sort.Sort
	sort.Sort(heroes)
	fmt.Println("排序后为\n")


	//看看排序后的顺序
	for _,v := range heroes{
		fmt.Println(v)
	}
}

