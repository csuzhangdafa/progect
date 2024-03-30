/*
map 是key-value数据结构，又称为字段或者关联数组。类似其它编程语言的集合，在编程中是经常使用到
基本语法：var map 变量名 map[keytype]valuetype
golang中的map，的 key 可以是很多种类型，比如 bool, 数字，string, 指针,channel, 
还可以是只包含前面几个类型的 接口, 结构体, 数组
valuetype的类型和key基本一样

map声明举例：
var a map[string]string
var a map[string]int
var a map[int]string
var a map[string]map[string]string
注意：声明是不会分配内存的，初始化需要make ，分配内存后才能赋值和使用

meke(type,size)

map[“key”]=value//如果 key 还没有，就是增加，如果 key 存在就是修改
说明：delete(map，“key”) ，delete是一个内置函数，如果key存在，就删除该key-value,如果key不存在，不操作，但是也不会报错

*/

package main
import (
	"fmt"
	_"sort"
)

func main(){
	//var map变量名map(关键字必须有)[keytype3eee]valuetype
	//1.map声明使用map前需要先make
	var a map[string]string
	a=make(map[string]string,10)//前面为map的数据类型，后面为分配的内存.key不能相同，如果重复了，则以最后这个key-value为准。value可以相同
	a["1"]="宋江"
	a["2"]="无用"
	a["1"]="武松"
	a["3"]="无用"
	fmt.Println(a)


	//map查询两个变量 value key
	val,ok :=a["1"]
	if ok{
		fmt.Println(val)
	}else{
		fmt.Println("无")
	}


    //2.make方式
	cities :=make(map[string]string)
	cities["1"]="北京"
	cities["2"]="天津"
	cities["3"]="上海"
	fmt.Println(cities)
	fmt.Println("cities, 有",len(cities),"对key-value")//map的长度


	//map遍历使用 for range  key value
	for k,v := range cities{
        fmt.Printf("k=%v,v=%v\n",k,v)
	}
	

	//3.直接声明赋值  var heros map[string]string=map[string]string
	heros :=map[string]string{
		"1":"松江",
		"2":"卢俊义",
		"3":"武松",
	}
	heros["4"]="林冲"    //增加，如果有值就是修改
	fmt.Println(heros)


	//案例存放三个学生的信息每个学生有name sex信息 map[string]map[string](后面value作为一个map存放学生信息)
	studentsmap :=make(map[string]map[string]string)
	studentsmap["1"] = make(map[string]string,3)//使用之前要make
	studentsmap["1"]["name"]="tom"
	studentsmap["1"]["sex"]="man"
	studentsmap["1"]["address"]="北京长安街"


	studentsmap["2"] = make(map[string]string,3)
	studentsmap["2"]["name"]="mary"
	studentsmap["2"]["sex"]="woman"
	studentsmap["2"]["address"]="上海"
	fmt.Println(studentsmap["2"]["address"])
	fmt.Println()
	fmt.Println(studentsmap)
	fmt.Println()


    //for range遍历这个较为复杂的结构两层循环 外层为value，内层为key
    for k1,v1 :=range studentsmap{
		fmt.Println("k1=",k1)
		for k2,v2 :=range v1{
		    fmt.Printf("\t k2=%v v2=%v\n",k2,v2)
		}
		fmt.Println()
	}


	//map增删改查map["key"]=value,如果没有key、就是增加，如果存在就是修改
	//删除delete(map,"key")delete为内置函数如果key存在就删除，如果不存在也不会报错
	delete(studentsmap,"1")
	fmt.Println(studentsmap)
	fmt.Println()


	//如果希望一次删除所有的key 1，遍历所有的key逐一删除，2，直接make一个新的空间
	studentsmap = make(map[string]map[string]string)
	fmt.Println(studentsmap)
	fmt.Println()
   

	//map切片用map来记录monster的name和age信息，要求可以动态增加
	var monsters []map[string]string  //声明一个map类型的切片
	monsters = make([]map[string]string , 2)//切片本身要make

	if monsters[0]==nil{
	    monsters[0]=make(map[string]string,2)//切片对应的map类型也要make
	    monsters[0]["name"]="牛魔王"
	    monsters[0]["age"]="500"
	}
	if monsters[1]==nil{
	    monsters[1]=make(map[string]string,2)//切片对应的map类型也要make
	    monsters[1]["name"]="玉兔精"
	    monsters[1]["age"]="500"
	}
	fmt.Println(monsters)


	//map类型slice使用append函数动态增加
	monster1 :=map[string]string{
		"name" :"新的妖怪",
		"age" :"200",
	}
	monsters = append(monsters,monster1)
	fmt.Println(monster1)


	//map排序
	map1 := make(map[int]int,10)
	map1[6]=10
	map1[2]=49
	map1[5]=95
	map1[3]=13
	map1[8]=35
	fmt.Println(map1)

	//现在已经默认按照key排序
	//按照map的key顺序进行排序输出
	//1.先将map的key放到切片中 2，对切片进行排序 3，遍历切片

}