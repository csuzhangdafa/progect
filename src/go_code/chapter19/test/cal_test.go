/*
工作中去确认一个函数或者模块的结果是否正确
1.传统方法：直接调用，查看输出结果是否正确。
缺点：1.放到主函数中去运行，需要修改主函数，如果项目正在运行可能会停止
    2.不利于管理，测试多个模块时不利于管理。
	3.引出单元测试
	1 确保每个函数是可运行，并且运行结果是正确的
    2 确保写出来的代码性能是好的，
    3 单元测试能及时的发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，
      而性能测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定
	  func TestXxx(*testing.T)
*/

package main
import(
	_"fmt"
	"testing"
)

func TestAddUpper(t *testing.T){
	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)//t结构体自带的方法
	}

	//如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确...")
}

//没有主函数依然可以运行，是因为调用了test框架以将_test.go引入主函数
//cd d:\go_project\pkg\src\go_code\chapter19
//go test