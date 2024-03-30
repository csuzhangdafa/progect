package service
import (
	"go_code/chapter15/customermanage/model"
)

//完成对Customer的操作
type CustomerService struct{
	customers []model.Customer
    //申明一个字段，表示当前切片含有多少个客户,此切片可以增加，此字段包含的是modle包中的customer结构体

	//该字段还可以作为新客户的id+1
	customerNum int
}


//编写一个函数可以返回*Customersercice
func NewCustomerService() *CustomerService{

	//为了能够看到有客户在切片中，我们初始化一个切片

	customerService := &CustomerService{}     //对结构体初始化
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"张三","男",20,"112","zs@qq")  //对modle中的customer结构体初始化
	customerService.customers = append(customerService.customers, customer)  //把初始化的customer结构体加入到切片中
	return customerService
}


//返回客户切片
func (this *CustomerService) List() []model.Customer{
	return this.customers
}


//添加客户到customer切片中
func (this *CustomerService) Add(customer model.Customer) bool{


	//确定一个分配id的规则,就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return true

}


//根据id删除客户
func (this *CustomerService) Delete (id int) bool {
	index := this.Findbyid(id)
	
	//如果index = -1 说明没有这个客户
	if index == -1{
		return false
	}

	//从切片中删除一个元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...)  //append的使用前一个不包括index，后一个到最后
	return true
}


//根据id查找客户在切片中对应的下标，如果没有该客户，返回-1
func (this *CustomerService) Findbyid(id int) int {

	index := -1
	//遍历this.customers 切片
	for i :=0; i<len(this.customers); i++{
		if this.customers[i].Id == id{

			//找到
			index = i
		}
	}

	return index
}