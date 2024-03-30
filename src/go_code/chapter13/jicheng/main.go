package main
import (
	"fmt"
)


//继承案例之编写一个学生考试系统


//将共有的方法也绑定到student
type Student struct{
    Name string
	Age int
    Score int 
}


type Pupil struct{
    Student  //嵌入student匿名结构体
}


type Graduate struct{
    Student  //嵌入student匿名结构体
}


//显示成绩
func (stu *Student) Showinfo(){
	fmt.Printf("学生名字=%v 年龄=%v 成绩=%v\n",stu.Name,stu.Age,stu.Score)
}


func (stu *Student) Setscore(score int){
	stu.Score = score
}


//特有的字段和方法都保留
func (p *Pupil) tesing(){
	fmt.Println("小学生正在考试")
}


func (p *Graduate) tesing(){
	fmt.Println("大学生正在考试")
}


//如果是大学生，高中生考试大部分代码相同，则会出现代码冗余
//通过使用嵌套匿名结构体的方式来实现继承特性


func main(){
    var pupil = &Pupil{}
	pupil.Student.Name = "tom"  //.Student可以简化
	pupil.Student.Age = 8
	pupil.tesing()
	pupil.Student.Setscore(70)
	pupil.Student.Showinfo()


	var graduate = &Graduate{}
	graduate.Student.Name = "marry"
	graduate.Student.Age = 18
	graduate.tesing()
	graduate.Student.Setscore(90)
	graduate.Student.Showinfo()
}
