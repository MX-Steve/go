package main

import (
	"fmt"
)

// 接口嵌套

type speaker interface {
	speak()
}
type mover interface {
	move()
}
type animal interface {
	speaker
	mover
}
type cat struct {
	name string
}

func (c cat) move() {
	fmt.Println("移动")
}
func (c cat) speak() {
	fmt.Println("说话")
}

// 类型断言
func justifyType(x interface{}) {
	// 类型断言：这个函数可以接收任意类型的参数
	// _, ok := x.(int)
	// if !ok {
	// 	fmt.Println("x不是一个int类型")
	// }else {
	// 	fmt.Println("x就是一个int类型")
	// }
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int, is %v\n", v)
	case bool:
		fmt.Printf("x is a bool,is %v\n", v)
	case struct{}:
		fmt.Printf("x is a struct{},is %v\n", v)
	case cat:
		fmt.Printf("x is a cat,is %v\n", v)
	default:
		fmt.Println("未知类型")
	}
}

func main() {
	// var x animal
	// x = cat{name:"花花"}
	// x.move()
	// x.speak()
	// var studentList = make([]map[string]interface{},0)
	// var studentInfo = make(map[string]interface{})
	// studentInfo["name"]="张三"
	// studentInfo["age"]=23
	// studentInfo["married"]=false
	// fmt.Println(studentInfo)
	// studentList = append(studentList,studentInfo)
	// fmt.Println(studentList)
	// var studentInfo2 = make(map[string]interface{})
	// studentInfo2["name"]="王五"
	// studentInfo2["age"]=20
	// studentInfo2["married"]=true
	// fmt.Println(studentInfo2)
	// studentList = append(studentList, studentInfo2)
	// fmt.Println(studentList)

	var x interface{}
	x = 100
	justifyType(x)
	justifyType("abc")
	justifyType(struct{}{})
	y := "行号"
	justifyType(&y)
}
