package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s Student) Study() string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}

func (s Student) Sleep() string {
	msg := "好好睡觉，快快长大"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method type:%s\n", methodType)
		var args = []reflect.Value{}
		v.Method(i).Call(args) // Call() 参数必须是 reflect.Value的切片类型
	}
}

func main() {
	var stu1 = Student{
		Name:  "李四",
		Score: 39,
	}
	printMethod(stu1)
}
