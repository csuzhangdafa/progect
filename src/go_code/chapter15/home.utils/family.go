package utils
import (
	"fmt"
)


/*家庭收支记账软件

1.项目开发流程说明
  1.需求分析    30%
  2.设计阶段    20%
  3.实现阶段    20%
  4.测试阶段    螺旋递增
  5.实施阶段
  6.维护阶段
2.项目需求说明
3.项目的界面
4.项目的代码实现

实现基本功能（先用面向过程，然后改成面向对象）
功能1.完成可以显示主菜单，并且可以退出
功能2.完成可以显示明细的功能
功能3.完成登记收入，支出的功能
*/


//把所有的字段封装到结构体，相应的方法与结构体进行绑定,最后提供一个工厂模式的构造方法
type Family struct{
	key string   //保存用户输入选项

	//声明一个变量控制循环
	loop bool

	//定义账户余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义变量，记录是否有收支的行为
	flag bool
	//收支的详情用details来记录
	details string

}


//编写一个工厂模式的构造方法返回一个*Family的实例，构造一个函数来访问私有属性，结构体首字母都是小写的，不能直接访问
//对结构体实例化，通过工厂函数实例化的实例也可以使用结构体相应的方法
/*工厂函数一般化定义
package mypackage

type MyStruct struct {
	// 结构体的字段定义
}

// NewMyStruct 是一个工厂函数，用于创建和初始化 MyStruct 结构体实例。
func NewMyStruct() *MyStruct {
	// 使用结构体字面量初始化字段
	return &MyStruct{
		// 初始化字段值
	}
}
*/

func NewFamily() *Family{
	return &Family{
		key : " ",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : " ",
		flag : false,
		details : "收支\t账户金额\t收支金额\t说明",
	}
}



//将显示明细写成一个方法
func (this *Family)Showdetails(){
	fmt.Println("\n.....................当前收支记录明细.................")
	if this.flag{
		fmt.Println(this.details)
	}else{
		fmt.Println("当前没有收支")
	}
}


//登记收入方法
func (this *Family)Income(){
	fmt.Println("本次收入金额")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v    \t%v     \t%v",this.balance,this.money,this.note)
	this.flag = true

}


//登记支出的方法
func (this *Family)Outcome(){
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)
	if this.money > this.balance{
		fmt.Println("支出的余额不足")
		return 
	}
	this.balance -= this.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v    \t%v     \t%v",this.balance,this.money,this.note)
	this.flag = true
}


//退出系统的方法
func (this *Family)Exit(){
	fmt.Println("确定要退出嘛y/n\n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y"|| choice == "n"{
                    break
				}
                fmt.Println("你输入的有误请输入正确的选项 y/n")
			}
			if choice == "y"{
				this.loop = false
			}
}


//给该结构体绑定相应的方法
//显示主菜单,主菜单因该是循环显示 
func (this *Family)Mainmenu(){
	for{
		fmt.Println(".....................家庭收支记账软件.................")
		fmt.Println("                     1.收支明细                      ")
		fmt.Println("                     2.登记收入                      ")
		fmt.Println("                     3.登记支出                      ")
		fmt.Println("                     4.退出软件                      ")
		fmt.Println("请选择(1-4)")
        fmt.Scanln(&this.key)


		switch this.key{
		case "1":
			this.Showdetails()
		case "2":
			this.Income()
		case "3":
			this.Outcome()
		case "4":
			this.Exit()
		default:
			fmt.Println("                   请输入正确的选项                ")
		}

		if !this.loop{
			break
		}
	}
    fmt.Println("                   你退出了软件的使用                        ")
}