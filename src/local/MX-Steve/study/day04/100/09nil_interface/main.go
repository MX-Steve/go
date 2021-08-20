package main

import (
	"fmt"
	"time"
)

func showType(x interface{}) {
	fmt.Printf("%T\n", x)
}

// type assert
func showType2(x interface{}) {
	// v, ok := x.(int)
	// if !ok {
	// 	fmt.Println("不是")
	// } else {
	// 	fmt.Println(v)
	// }
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string.value is %v\n", v)
	case int:
		fmt.Printf("x is a int. value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool. value is %v\n", v)
	case struct{}:
		fmt.Printf("x is a bool. value is %v\n", v)
	default:
		fmt.Println("x's type is not known")
	}
}

// nil interface
func main() {
	// var x interface{}
	// x = 100
	// x = "string"
	// x = true
	// x = false
	// x = struct{ name string }{name: "xx"}
	// fmt.Println(x)
	// showType(100)
	// showType("stri")
	testIt := true
	if testIt {
		showType2("xss")
		showType2(18)
		showType2(true)
		showType2(struct{}{})
	} else {
		showType(time.Second)
		var stuInfo = make(map[string]interface{}, 100)
		stuInfo["lisi"] = 100
		stuInfo["zhangsan"] = true
		stuInfo["now"] = time.Now()
		fmt.Println(stuInfo)
	}
}
