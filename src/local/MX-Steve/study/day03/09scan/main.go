package main 

import (
	"fmt"
)

// 从终端标准输入获取用户输入内容
func  main(){
	var (
		name string 
		age int 
		married bool
	)
	// fmt.Scan(&name,&age,&married)
	// fmt.Scanf("name:%s age:%d married:%T\n", &name, &age, &married)
	fmt.Println("name age married")
	fmt.Scanln(&name,&age,&married)
	fmt.Println(name,age,married)

}