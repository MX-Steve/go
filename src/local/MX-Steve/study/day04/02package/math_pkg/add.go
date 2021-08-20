package math_pkg

import (
	"fmt"
)

// Add 加法
func Add(x,y int) int{
	return x + y
}

func init(){
	fmt.Println("math_pkg 中的init")
}