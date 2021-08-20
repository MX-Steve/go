package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "张三",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	// 通过结构体遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	if scoreFiled, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d,type:%v,json tag:%v\n", scoreFiled.Name, scoreFiled.Index, scoreFiled.Type, scoreFiled.Tag.Get("json"))
	}
}
