package main 

import (
	"fmt"
)

func main(){
	a := make([]int, 3)
	a[0]=1
	a[1]=2
	a[2]=3
	fmt.Println(a)
	fmt.Printf("a的指针：%p\n",a)
	b := make([]int,3)
	copy(b,a)
	fmt.Println(b)
	fmt.Printf("b的指针：%p\n",b)
	c := a[:]
	fmt.Println(c)
	fmt.Printf("c的指针：%p\n",c)
	a = append(a[:1],a[2:]...)
	fmt.Println(a)
}