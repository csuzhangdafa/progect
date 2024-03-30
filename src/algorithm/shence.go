package main

import (
	"fmt"
)

/*
给定一个数组arr，表示连续n天的股价，数组下标表示第几天
指标x：任意两天股价之和-此两天的间隔
第三天，10
第9天，30
指标：x=10+30-（9-3）=34
返回最大指标x
时间复杂度o(n) 
*/
func main() {
	fmt.Println("main")
	
}
