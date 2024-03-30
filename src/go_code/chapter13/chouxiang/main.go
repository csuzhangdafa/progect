package main
import (
	"fmt"
)


//我们在定义一个结构体的时候，实际上是把一类事物的共有属性（字段）和行为（方法）提取出来
//形成一个物理模板，这种研究问题的方法称为抽象


type Account struct{
	AccountNo string  //账户
	Pwd string        //密码
	Balance float64   //余额
}


//方法1.存款
func (account *Account) Deposite(money float64,pwd string){ //存款数目和密码
    
	//比对密码
	if pwd !=account.Pwd{
		fmt.Println("你输入的密码不正确\n")
		return
	}
	//看看存款金额是否合理
	if money <= 0{
		fmt.Println("你输入的金额不正确\n")
		return
	}
	account.Balance += money
	fmt.Println("存款成功\n")
}


//取款
func (account *Account) WithDraw(money float64,pwd string){ 
    
	//比对密码
	if pwd !=account.Pwd{
		fmt.Println("你输入的密码不正确\n")
		return
	}
	//看看存款金额是否合理
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确\n")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功\n")
}


//查询余额
func (account *Account) Qurry(pwd string){ 
    
	//比对密码
	if pwd !=account.Pwd{
		fmt.Println("你输入的密码不正确\n")
		return
	}
	fmt.Printf("你的账号余额=%v\n",account.Balance)
}

func main(){
    //测试
	account :=&Account{
		AccountNo : "工商银行1111",
	    Pwd : "666666",     
	    Balance : 100.0, 
	}


	//可以做的更灵活
	account.Qurry("666666")
	account.Deposite(200.0,"666666")
	account.Qurry("666666")
	account.WithDraw(150.0,"666666")
	account.Qurry("666666")
}





