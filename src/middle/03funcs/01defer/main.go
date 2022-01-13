package main

import "fmt"

func d1() {
	var whatever [5]struct{}
	for i := range whatever {
		defer fmt.Println(i)
	}
}
func d2() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}
func Close(t Test) {
	t.Close()
}
func d3() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Close()
	}
}
func d4() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer Close(t)
	}
}
func d5() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		var t2 = t
		defer t2.Close()
	}
}
func test(x int) {
	defer fmt.Println("a")
	fmt.Println("b")
	defer func() {
		fmt.Println(100 / x)
	}()
	defer fmt.Println("c")
}
func test2() {
	x, y := 10, 20
	defer func(i int) {
		fmt.Println("defer: ", i, y)
	}(x)
	x += 10
	y += 100
	fmt.Println("x=", x, " y=", y)
}
func main() {
	test2()
}
