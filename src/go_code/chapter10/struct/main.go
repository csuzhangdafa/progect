/*type 结构体名称 struct{
	field1 type
	field2 type
}
*/

package main
import (
	"fmt"
)


//结构体入门
type cat struct{
	name string //field type首字符大写表示可以在别的包中引用,小写为私有只有在本包中使用
	age int
	color string
}


//如果结构体的字段类型为指针，map slice，默认的值都为nil还没有分配空间，如果要使用先make
type person struct{
	name string
	age int
	scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}


func main(){
	//结构体基本入门
    var cat1 cat
	cat1.name = "小白"
	cat1.age = 3
	cat1.color = "白色"
	fmt.Println("猫猫的信息如下\n")
	fmt.Println("name=",cat1.name)
	fmt.Println("age=",cat1.age)
	fmt.Println("color=",cat1.color)


	//slice map 使用前make
	var p1 person
	p1.slice = make([]int,10)
	p1.slice[0] = 100
	p1.map1 = make(map[string]string)
	p1.map1["1"] = "上海"
	p1.map1["2"] = "北京"
	fmt.Println(p1)


	//不同的结构体变量字段是独立的，互不影响一个更改不影响另一个，结构体是值类型


	//创建结构体变量和使用结构体字段的方式
	//1. 直接申明 var person Person


	//2.var person Person = Person{}
	p2 := person{}//person{"mary",20}
	p2.name = "tom"
	p2.age = 18
	fmt.Println(p2)


	//3.var person*Person = new(Person)
	p3 := new (person)  //var p3 *person = new(person) 因为p3为一个指针，标准的赋值方式为(*p3).name = ""
	p3.name = "smith"   //go 设计者在底层做处理优化等价于(*p3).name = "smith"
	(*p3).age = 19
	fmt.Println(*p3) 
	
	
	//4.var person*person = &person{},与3使用方法相同
	var p4*person = &person{} 
	p4.name = "mary"   
	(*p4).age = 20
	fmt.Println(*p4) 


	//创建结构体变量指定字段的值
    /*type stu struct{
		age int，
		name string，
	} 
	1. var stu1 = stu{"小明",19}

    2. var stu3 = stu{
		name : "jack",
		age : 20,
	  }


	//使用细节1.结构体中的所有字段在内存中都是连续分配的
	//2.结构体是用户单独定义的字段，和其他的类型进行转换时要有完全相同的字段，类型，名称，个数完全相同
	//3.结构体用type重新定义（相当于取别名）,golang认为是新的数据类型，不可以直接赋值，但是之间可以相互强转
	//4.struct的每一个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化*/
}
    
