package main 

import (
	"fmt"
)

// 实现接口时使用指针接受者和使用值接受者的区别
type animal interface{
	speak()
	move()
}

type cat struct {
	name string 
}

// cat 类型要实现animal的接口
// 方法一：使用值接受者
/*
func (c cat) speak(){
	fmt.Println("咪咪咪")
}
func (c cat) move(){
	fmt.Println("runrunrun")
}
*/
// 方法二：使用指针接受者
func (c *cat) speak(){
	fmt.Println("咪咪咪")
}
func (c *cat) move(){
	fmt.Println("runrunrun")
}

func main(){
	var x animal 
	// hh := cat{"花花"}
	// x = hh // 实现animal接口的是 *cat 类型。hh此时是cat 类型[值类型]
	tom := &cat{"tom"} // *cat
	fmt.Println(x) 
	x = tom 
	tom.speak() //(*tom).speak()
	tom.move() //(*tom).move()
	fmt.Println(x)
}