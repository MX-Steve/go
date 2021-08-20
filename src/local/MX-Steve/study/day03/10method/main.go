package main

import (
	"fmt"
)

// 方法就是某个具体的类型才能调用的函数
// 函数是谁都可以调用的
type people struct {
	name string 
	gender string 
}

// 以下情况使用指针
// 需要修改接受者的值
// 接受者是拷贝代价比较大的对象
// 一般情况都使用指针
func (p *people)dream(){
	p.gender="女"
	fmt.Printf("%s(%s)的梦想是不用上班也有钱！\n",p.name,p.gender)
}

func main(){
	var lisi = people{"李四","男"}
	lisi.dream()
	fmt.Println(lisi)
}