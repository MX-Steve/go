package main

import (
	"fmt"
)

// 空接口

func main(){
	var x interface{}
	x = "hello"
	fmt.Println(x)
	x = 10
	fmt.Println(x)
	x = struct{name string}{name:"张三"}
	fmt.Println(x)
}