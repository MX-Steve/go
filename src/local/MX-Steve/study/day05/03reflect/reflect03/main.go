package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改值

func modifyValue(x interface{}) {
	v := reflect.ValueOf(x) // reflect.Value
	kind := v.Kind()
	fmt.Println(kind)
	if kind == reflect.Ptr {
		// 传入的是一个指针
		// 取到指针对应的值进行修改, 必须使用 Elem() 取出值进行修改
		v.Elem().SetInt(12)
	}
}

func main() {
	var x int64 = 100
	modifyValue(&x)
	fmt.Println(x)
}
