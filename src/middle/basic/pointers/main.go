package main

import "fmt"

// 指针地址，指针类型，指针取值
// & 取地址， * 根据地址取值
func f1() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr: %p\n", a, &a)
	fmt.Printf("b: %p type: %T\n", b, b)
	fmt.Println(*b)
}
func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 200
}
func f2() {
	var p *string
	fmt.Println(p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}
func f3() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)
	var c *string
	c = new(string)
	*c = "abc"
	fmt.Println(*c)
}
func f4() {
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
}

func main() {
	f4()
	f3()
	f2()
	a := 10
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
	f1()
}
