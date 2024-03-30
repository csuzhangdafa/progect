package main
import(
	"fmt"
	"go_project/pkg/src/go_code/chapter13/model"
)


//面向对象编程三大特征，封装，继承，多态，golang也存在，只是设计和传统的oop语言不太相同


/*封装就是把抽象出来的字段和对字段的操作封装在一起，数据被保护在其内部，
程序的其他包只有通过被授权的操作（方法）才能对字段进行操作
将结构体字段，属性小写，在其他包中不可以使用，类似于工厂模式函数


提供一个首字母大写的Set方法类似于其他函数的public，对于属性判断并赋值
func (var 结构体类型名) SetXxx(参数列表)(返回值列表){
	加入数据验证的业务逻辑
	var.字段 = 参数
}


提供一个首字母大写的Get方法类似于其它语言的public，对于获取属性的值
func (var 结构体类型名) GetXxx{
	return var.字段
}*/

func main(){
    p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name,"age=",p.GetAge(),"Sal=",p.GetSal())

}
