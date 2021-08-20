package main 

import (
	"fmt"
)

// 结构体是一个值类型
type  student struct{
	name  string
	age   int8
}


func main(){
	var stu1 = student{
		name: "李四",
		age: 28,
	}
	var stu2 = stu1 // 将结构体 stu1 的值完整的复制一份给 stu2
	stu2.name="张三"
	fmt.Println(stu1)
	fmt.Println(stu2)
	stu3 := &stu1
	stu3.name="王五"
	fmt.Println(stu1)
	fmt.Println(stu3)
	fmt.Printf("%T\n",stu3)
}