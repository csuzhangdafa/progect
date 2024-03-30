package main
import (
	"fmt"
	"encoding/json"
)


type Monster struct{
    Name string `json:"monstername"`  //指定序列化的key值,反射机制
	Age int `json:"monsterage"`
	Birthdy string
	Sal float64
	Skill string
}


//将结构体序列化
func teststruct(){
	monster := Monster{
		Name : "牛魔王",  
	    Age : 34,
	    Birthdy : "2011/11",
	    Sal : 8000.3,
	    Skill : "牛魔全",
	}
	//将monster序列化
	data,err := json.Marshal(&monster)
	if err !=nil{
		fmt.Printf("序列化错误%v\n",err)
	}
	fmt.Printf("%v\n",string(data))
}


//将map序列化
func testmap(){
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 20
	a["address"] = "洪崖洞"

	//将monster序列化
	data,err := json.Marshal(a)
	if err !=nil{
		fmt.Printf("序列化错误%v\n",err)
	}
	fmt.Printf("%v\n",string(data))
}


//将切片的序列化
func testslice(){
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 20
	m1["address"] = "北京"
	slice = append(slice , m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 19
	m2["address"] = "南京"
	slice = append(slice , m2)


	//对切片序列化
	data,err := json.Marshal(slice)
	if err !=nil{
		fmt.Printf("序列化错误%v\n",err)
	}
	fmt.Printf("%v\n",string(data))

}


//对基本数据类型序列化,意义不大
func testfloat(){
	var num1 float64 = 235.12
	data,err := json.Marshal(num1)
	if err !=nil{
		fmt.Printf("序列化错误%v\n",err)
	}
	fmt.Printf("%v\n",string(data))

}


//json是一种轻量级的数据交换模式，也是主流的数据格式,任何数据类型都可以转成json格式
//[{"key1":val1,"key2":"val2","key3":"val3","key4":[val4,val5]},      ]


func main(){

	//数据格式序列化,首字母大写

	teststruct()
	testmap()
	testslice()
	testfloat()
}