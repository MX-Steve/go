package main

import (
	"fmt"
)

func main(){
	// var a = [3]int{1,2,3}
	// var b = []int{1,2,3}
	// fmt.Println(a,b)
	// fmt.Printf("a:%T; b:%T\n",a,b)
	// fmt.Println(b[1])

	// var c []int 
	// c =a[:]
	// fmt.Printf("c:%T\n",c)
	// fmt.Println(c)
	// var d []int 
	// d = a[:2]
	// fmt.Println(d)
	// e := a[1:]
	// fmt.Println(e)

	// fmt.Println(len(b))
	// x := [...]string{"北京","上海","广州","深圳","成都","西安","杭州"}
	// y := x[1:4]
	// fmt.Println(y)
	// fmt.Println(cap(y))

	// a := []string{"北京","上海","广州","深圳"}
	// a = append(a[:1],a[2:]...)
	// fmt.Println(a)

	// a := [...]int{1,3,5,7,9,11,13}
	// b := a[:]
	// b[0]=100
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Printf("b:%p\n",b)

	a := [...]int{1,3,5,7,9,11,13}
	c := a[:5]	
	fmt.Println(c)
	fmt.Printf("c:%p\n",c)
	d := c[:6]
	fmt.Println(d)
	fmt.Printf("d:%p\n",d)
	c[0]=100
	fmt.Println(c)
	fmt.Println(d)

}