package main
import (
	"fmt"
	"go_code/chapter15/customermanage/service"
	"go_code/chapter15/customermanage/model"
)


type customerView struct{

	//定义必要的字段
	//接受用户的输入，1-5
	key string
	//表示是否循环显示菜单，前一个例子用过此方法
	loop bool
	//增加一个字段, 因为要调用customerservice,所以要增加一个字段
	customerService *service.CustomerService
}


//显示所有的客户信息调用service中的List方法
func(this *customerView) list(){

	//获取当前多有的客户信息，信息存在于切片中。this.customerservice是指customerview中的customerservice变量
	//此变量类型又是调用service包中的CustomerService结构体
	//在主函数中已经对customerService进行了初始化，因此可以调用相关方法
	//this->customerService->service.CustomerService->NewcustomerService（初始化）->list（返回客户信息）
	customers := this.customerService.List()
    //customers中为所有的客户，为一个切片


	//显示
	fmt.Println(".....................客户列表.......................")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0 ; i<len(customers);i++{
		fmt.Println(customers[i].GetInfo())  //此方法为modle包中的GetInfo方法
	} 
	fmt.Println("...................客户列表完成.....................\n\n")
}



//调用Add方法
func(this *customerView) add(){
    fmt.Println("....................添加客户..........................")
    fmt.Println("姓名:")
	name :=" "
	fmt.Scanln(&name)
    fmt.Println("性别:")
	gender :=" "
	fmt.Scanln(&gender)
    fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
    fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
    fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)


	//构建一个新的Customer实例 id需要系统分配
	customer :=model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer){
		fmt.Println("....................添加完成.....................")
	}else{
		fmt.Println("....................添加失败.....................")
	}
}


//得到用户输入的id,删除该id对应的客户
func(this *customerView) delete(){

	fmt.Println(".....................删除客户.....................")
	fmt.Println("请选择待删除的客户编号(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return
	}
	fmt.Println("确认是否删除(Y/N)")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y"||choice =="y"{

		if this.customerService.Delete(id){
            fmt.Println("....................删除完成..................")
		}else{
			fmt.Println("....................删除失败,请重新输入........")
		}
	}
	
}


//显示主菜单
func(this *customerView) mainMenu(){

	for{
		fmt.Println("...................客户信息管理软件.................")
		fmt.Println("                   1.添加客户                     ")
		fmt.Println("                   2.修改客户  ")
		fmt.Println("                   3.删除客户 ")
		fmt.Println("                   4.客户列表 ")
		fmt.Println("                   5.退出 ")
		fmt.Println("                   请选择(1-5)")


		fmt.Scanln(&this.key)
		switch this.key{
	    case "1": 
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
		if !this.loop{
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统")
}



func main(){
	//在主函数中，创建一个customerview实例，并运行显示主菜单，结构体实例化以后才可以调用相关方法
	customerView := customerView{
		key : " ",
		loop : true,
	}

	//对customerview结构体里面的customerservice初始化，初始化方法是调用service包内的NewCustomerService()工厂函数
	//在此函数内，已经完成了customersevice的初始化。
	customerView.customerService = service.NewCustomerService()

	 //显示主菜单
	customerView.mainMenu() 
}