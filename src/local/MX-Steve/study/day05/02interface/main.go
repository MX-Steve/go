package main

import (
	"fmt"
)

func justifyType(x interface{}){
	switch v := x.(type){
	case string:
		fmt.Printf("x is a string, value is %v\n",v)
	case int:
		fmt.Printf("x is a int, value is %v\n",v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n",v)
	default:
		fmt.Println("unsupport type!")
	}
}

func main(){
	var x interface{}
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1
	x = a // <int64, 100>
	x = b // <int32, 10>
	x = c // <int8, 1>
	x = 12.34 // <float64, 12.34>
	x = false // <bool, false>
	// fmt.Println(x)
	// value,ok := x.(int)
	// if ok{
	// 	fmt.Printf("x是int类型，值是%v\n",value)
	// }else {
	// 	fmt.Println("x存的不是int类型")
	// }
	// 类型断言（猜）
	value, ok := x.(float64)
	fmt.Printf("ok:%t; value:%v; value type:%T\n",ok,value,value)
	justifyType(23)
	justifyType("abc")
	justifyType(true)
	justifyType(map[string]int{"age":23})
}