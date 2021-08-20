package main

import (
	"fmt"
	"reflect"
)

// 反射的TypeOf()
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	// fmt.Printf("type:%v\n", t)
	fmt.Printf("name:%v; kind:%v\n", t.Name(), t.Kind())
}

type cat struct {
	name string
}

type person struct {
	name string
	age  uint8
}

func main() {
	// reflectType(100)
	// reflectType("abc")
	// reflectType(false)
	// reflectType([3]int{1, 2, 3})
	// reflectType(map[string]int{})

	// 反射拿到类型
	var c1 = cat{
		name: "花花",
	}
	var p1 = person{
		name: "李四",
		age:  18,
	}
	reflectType(c1)
	reflectType(p1)
	var n = 100
	reflectType(&n)
	var i int32 = 100
	reflectType(i)
	reflectType(&i)
	// 应用类型的 t.Name() 为空
}
