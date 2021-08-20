package main

import (
	"fmt"

	_ "local/MX-Steve/study/day04/02package/math_pkg"
)

// init 实例
var today = "星期三"

const Week = 7

type Student struct {
	Name string
}

// init() 包被导入时会自动执行（多用来做初始化操作）
func init() {
	fmt.Println("这是一个初始化init 操作")
	fmt.Println(today)
}

func main() {
	fmt.Println("main 函数")
	fmt.Println(Week)
	stu1 := &Student{
		Name: "lisi",
	}
	fmt.Println(stu1)
}
