package main 

import (
	"fmt"
)

// 只要一个类型实现了 wash 和 dry ，我们就称这个类型实现了洗衣机接口
type xiyiji interface {
	wash()
	dry()
}

type Haier struct {
	name string
	price float64
}

type Xte struct{
	name string
	price float64
}

func main(){
	var haier xiyiji
	fmt.Println(haier)
}