package main 

import (
	"fmt"
)

// 结构体内嵌模拟“继承”
type animal struct{
	name string 
}
func (a *animal)move(){
	fmt.Printf("%s 会动~\n",a.name)
}
type dog struct {
	feet int 
	animal 
}

func (d *dog)wangwang(){
	fmt.Printf("%s 在叫汪汪汪\n", d.name)
}

func main(){
	var a = dog{
		feet: 4,
		animal: animal{
			name: "旺财",
		},
	}
	a.move()
	a.wangwang()
}