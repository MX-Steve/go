package main

import (
	"fmt"
)

func main(){
	// mapSlice := make([]map[string]int,0,10)
	// fmt.Println(mapSlice)
	// mapSlice = append(mapSlice,map[string]int{"aaa": 10})
	// mapSlice = append(mapSlice,map[string]int{"bbb": 20})
	// fmt.Println(mapSlice)

	// mapSlice := make([]map[string]string, 3)
	// for i,v := range mapSlice{
	// 	fmt.Println(i,v)
	// }
	// fmt.Println("after init")
	// mapSlice[0] = make(map[string]string, 10)
	// mapSlice[0]["name"] = "张三"
	// mapSlice[0]["address"] = "上海"
	// for i,v := range mapSlice{
	// 	fmt.Println(i,v)
	// }

	scoreMap := make(map[string]int)
	scoreMap["zhangsan"] = 19
	scoreMap["lisi"] = 29
	v,ok := scoreMap["wangwu"]
	if ok{
		fmt.Println(v)
	}else {
		fmt.Println("no this person")
	}
}