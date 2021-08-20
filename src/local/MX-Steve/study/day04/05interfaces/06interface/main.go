package main

import (
	"fmt"
)

type Cat struct{}

func (c Cat)Say() string{return "喵喵喵"}

type Dog struct{}

func (d Dog)Say() string{return "汪汪汪"}

// 接口里的元素是一些方法名，包括函数名，参数，返回值
type Sayer interface {
	Say() string 
}



func main(){
	// c := Cat{}
	// fmt.Println("猫：",c.Say())
	// d := Dog{}
	// fmt.Println("狗：",d.Say())
	var animalList []Sayer
	c := Cat{}
	d := Dog{}
	animalList = append(animalList,c, d)
	fmt.Println(animalList)
}