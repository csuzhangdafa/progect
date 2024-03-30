package main
import(
	"fmt"
	_"math/rand"
	_"time"
) 


func ta(level int)(){ 
	for i := 1 ; i<=level ; i++{
		for k := 1 ; k<=level-i ; k++{
			fmt.Print(" ")
		}
		for j := 1 ; j<=2*i-1 ; j++{
				fmt.Print("*")
		}
		fmt.Println()
	}
}

func main(){
	var i int
	fmt.Println("请输入你的金字塔层数")
	fmt.Scanln(&i)
	ta(i)
	}
	
	/*for i := 1 ; i <= 9 ; i++{
		for j:= 1; j<=i ; j++{
			fmt.Printf("%v * %v = %v \t",j,i,j*i)
		}
		fmt.Println()
	}
	rand.Seed(time.Now().Unix())
	var count int = 0
	for {
		n := rand.Intn(100)+1
		count++
		if n == 99{
			break
	    }
    }
	fmt.Println(count)

	/*用户名，密码
	var name string 
	var pass string 
	var chance int = 3
	for i := 1 ; i<=3 ; i++{
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pass)
		if name == "张无忌"||pass == "888"{
			fmt.Println("登陆成功")
			break
		}else{
			chance--
			fmt.Printf("输入错误你还有%v次机会\n",chance)
		}
		if chance == 0{
			break
		}
	}*/
