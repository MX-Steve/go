package main

import (
	"fmt"
	"math"
)

func test(x, y int, z string) (int, string) {
	n := x + y
	return n, z
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func test2(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}
func main() {
	n, z := test(5, 10, "abc")
	fmt.Println(n, z)
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	fmt.Println(s2)
	fmt.Println(test2("sum: %d", 1, 2, 3, 4, 5))
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(9))
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	println(d.fn())
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello World!" }
	fmt.Println((<-fc)())
}
