package model
import(
	"fmt"
)

//声明一个customer结构体，表示客户

type Customer struct{
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Emile string
}


//客户信息会在cuntomerservice中使用因此，编写一个工厂模式，返回一个customer的实例
func NewCustomer(id int, name string ,gender string , age int , phone string  ,emile string)Customer{
	return Customer{
		Id : id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Emile : emile,
	}
}


//创建一个不带id的customer实例方法
func NewCustomer2(name string ,gender string , age int , phone string  ,emile string)Customer{
	return Customer{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Emile : emile,
	}
}


//返回用户的信息
func (this Customer) GetInfo() string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Emile)
	return info
}