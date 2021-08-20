package main

import "fmt"

var a = 10

func testGlobal() {
	a = 100
	fmt.Println(a)
}
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}
func main() {
	// fmt.Println(calc(4, 5, add))
	// fmt.Println(calc(4, 5, sub))
	// testGlobal()
	// fmt.Println(a)
	// fmt.Println(add(2, 3))
	// func() {
	// 	fmt.Println("hello world")
	// }()
	// closure
	ret := f1(10)
	fmt.Printf("type:%T\n", ret)
	fmt.Printf("num:%d\n", ret(5))
	fmt.Printf("num:%d\n", ret(6))
	fmt.Printf("num:%d\n", ret(7))
	fmt.Printf("num:%d\n", ret(8))
}

func f1(num int) func(x int) int {
	return func(x int) int {
		return num + x
	}
}
