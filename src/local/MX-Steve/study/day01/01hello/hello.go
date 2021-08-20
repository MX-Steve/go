package main

import "fmt"

func main() {
	var s string = "xsab"
	fmt.Println("Hello, kat.")
	fmt.Printf("hi,%s", "lihan\n")
	fmt.Println(s)
	name := "li汗"
	fmt.Println(name)
	var (
		a string
		b int
		c bool
	)
	a = "steve"
	b = 2
	c = true
	fmt.Println(a, b, c)
	// anonymous
	u, _ := foo()
	fmt.Println(u)
	var name2 = "Alex"
	fmt.Println(name2)
}

func foo() (string, int) {
	return "lihan", 20
}
