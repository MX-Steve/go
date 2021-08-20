package main 

import (
	"fmt"
)

func add(a,b int) int {
	return a + b
}

func add2(a,b int) (ret int) {
	ret =  a + b
	return
}

func add3(a int , b ...int) int {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("b的类型 %T",b)
	sum := a 
	for _,v :=range b{
		sum += v
	}
	return sum
}

func map1(){
	m1 := make(map[string][]string,3)
	fmt.Println("after init")
	fmt.Println(m1)
	key := "中国"
	value := make([]string,0,3)
	value = append(value,"北京","上海")
	m1[key]=value
	fmt.Println(m1)
}

func defer1(){
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end.")
}

var a = 100

func moda(){
	a = 200
	fmt.Println("in moda: ",a)
}
func moda2(){
	a := 200
	fmt.Println("in moda2: ",a)
}

func main(){
	i := add(1,2)
	fmt.Println(i)
	j := add2(20,30)
	fmt.Println(j)
	k := add3(1,22,3)
	fmt.Println(k)
	map1()
	defer1()
	// moda()
	// fmt.Println("out of moda: ",a)
	moda2()
	fmt.Println("out of moda2: ",a)
}