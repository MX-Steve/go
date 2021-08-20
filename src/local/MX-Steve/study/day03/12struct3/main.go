package main 

import (
	"fmt"
)

// 结构体的嵌套
type address struct {
	province string
	city string
}
type student struct {
	name string 
	age int
	addr address  // 嵌套了别的结构体
}

func main(){
	var stu1 = student{
		name: "李四",
		age: 28,
		addr: address{
			province: "江苏",
			city: "南京",
		},
	}
	fmt.Println(stu1)
	fmt.Println(stu1.addr.province)
}