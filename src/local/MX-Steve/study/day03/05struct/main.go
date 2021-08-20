package main 

import (
	"fmt"
)

// 结构体
// 创造一个新的类型用 type

type student struct {
	name string 
	age int 
	gender string 
	hobby []string
}


func main(){
	// 实例化方法1
	var zhangsan = student{
		name : "张三",
		age : 19,
		gender : "man",
		hobby : []string{"篮球","足球"},
	}
	fmt.Println(zhangsan)
	fmt.Println(zhangsan.name)
	fmt.Println(zhangsan.age)
	fmt.Println(zhangsan.gender)
	fmt.Println(zhangsan.hobby)
	
	// struct 是值类型的
	var wangzhan = student{}
	fmt.Println(wangzhan)
	wangzhan.age = 20
	wangzhan.gender = "female"
	fmt.Println(wangzhan)

	// 实例化方法2 new(T), T: 表示类型结构体
	var lisi = new(student)
	fmt.Println(lisi)
	lisi.name = "李四"
	lisi.age = 20
	lisi.gender = "male"
	lisi.hobby = []string{"足球","乒乓球"}
	fmt.Println(*lisi)

	// 实例化方法3
	var wangwu = &student{}
	fmt.Println(wangwu)

	// 结构体初始化
	// 只填值初始化
	var stu1 = student{"学生",18,"male",[]string{"男人","女人"}}
	fmt.Println(stu1)
	fmt.Println(stu1.name)
	// 键值对初始化
	var stu2 = &student{
		name : "学生2",
	}
	fmt.Println(stu2)
}