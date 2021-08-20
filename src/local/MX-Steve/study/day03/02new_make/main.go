package main 

import (
	"fmt"
)

func main(){
	// var a *int // a 是一个 int 类型的指针 , 错误写法
	var a = new(int) // 得到一个 int类型的指针
	// var b *string // 错误
	// var c *[3]int  // 错误
	*a = 10
	fmt.Println(a)
	fmt.Println(*a)
	// new 是用来初始化值类型指针的
	var c = new([3]int)
	fmt.Println(c)
	c[0] = 1
	c[1] = 2
	c[2] = 3
	fmt.Println(*c)

	// make 是用来初始化slice，map，channel 类型指针的
}