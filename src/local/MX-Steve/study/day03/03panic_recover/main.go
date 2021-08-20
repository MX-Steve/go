package main 

import (
	"fmt"
)

func main(){
	// panic 错误
	defer func(){
		// recover, 尝试从当前的异常恢复过来
		err := recover()
		fmt.Println("recover抓到了 panic 异常: ",err)
	}()
	var a[]int
	a[0] = 100 // panic
	fmt.Println("这是main函数")
}