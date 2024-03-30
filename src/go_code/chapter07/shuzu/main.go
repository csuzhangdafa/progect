/*
数组的初始化方式
var num1 [3]int = [3]int{1.2.3}

var num2 = [3]int{5.6.7}

var num3 = [...]int{8,9,10}

for index，value :=range array01{
	//for-range遍历方式
}

*/



package main 
import (
	"fmt"
)

func main(){

	//range便利i为下标，v为元素
	var herous [3]string = [3]string{"松江","无用","卢俊义"}//一维数组初始化

	//一维数组遍历
	for i ,v:=range herous{
		fmt.Printf("i=%v v=%v\n",i,v)
	}
	for _ ,v:=range herous{
		fmt.Printf("元素的值为=%v\n",v)
	}


	arr :=[...]string{"wer","qwe","qwe","qfa"}
	fmt.Println(arr)

	//一维数组循环赋值
	var mychars[26] byte
	for i :=0;i<26;i++{
		mychars[i]='A'+byte(i)
	}

	for i :=0;i<26;i++{
		fmt.Printf("%c \n",mychars[i])
	}


	//二维数组初始化赋值
	var arr1[2][3] int = [2][3]int{{1,2,3},{4,5,6}} //第一种方式
	var arr2[2][3] int = [...][3]int{{2,3,4},{4,5,6}}//第二种方式
	var arr3 = [2][3]int{{5,6,7},{8,9,6}}//第三种方式
	var arr4 = [...][3]int{{8,1,7},{8,9,6}}//第四种方式

	fmt.Println(arr1)
	fmt.Println()
	fmt.Println(arr2)
	fmt.Println()
	fmt.Println(arr3)
	fmt.Println()
	fmt.Println(arr4)
	fmt.Println()


	//二维数组遍历
	//1.用for循环遍历,两层for循环
	var arr5 = [2][3]int{{3,6,7},{4,9,6}}
	for i:=0;i<len(arr5);i++{
		for j:=0;j<len(arr5[i]);j++{
		    fmt.Printf("%v\t",arr5[i][j])
		}
	}
	fmt.Println()


	//使用for—range来遍历二维数组,i为横标，v为一维数组
	for i,v:=range arr5 {
		for j,v2 :=range v{

			//两层循环遍历
		    fmt.Printf("arr5[%v][%v]=%v\t",i,j,v2)//fmt.Printf("i=%v,v=%v\n",i,v)一层循环时可用
		}
		fmt.Println()
	}


	//二维数组应用案例,存三个班每个班五明学生的成绩
	var scores[3][5]float64
	for i:=0;i<len(scores);i++{
		for j:=0;j<len(scores[i]);j++{
			fmt.Printf("请输入第%d班第%d学生的成绩\n",i+1,j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	totalsum:=0.0
	for i:=0;i<len(scores);i++{
		sum :=0.0
		for j:=0;j<len(scores[i]);j++{
			sum+=scores[i][j]
		}
		totalsum +=sum
		fmt.Printf("第%d班级的总分为%v,平均分为%v\n",i+1,sum,sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分为%v,平均分为%v\n",totalsum,totalsum/15)
}