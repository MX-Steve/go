package main 

import (
	"fmt"
)

func main(){
	var x interface{}
	s := "hello 时间"
	x = s 
	fmt.Printf("类型：%T; 值: %v\n",x,x)
	i := 10
	x = i 
	fmt.Printf("类型：%T；值：%v\n",x,x)
	j := true
	x = j 
	fmt.Printf("类型：%T；值：%v\n",x,x)
}