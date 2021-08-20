package main

import (
	"fmt"

	_ "local/MX-Steve/study/day04/00total/01packages/test/hello"

	m "local/MX-Steve/study/day04/00total/01packages/math_pkg"
)

func main() {
	fmt.Println("xxx")
	result := m.Add(4, 5)
	fmt.Println(result)
	stu := m.Student{Name: "lisi", Age: 18}
	fmt.Println(stu)
	fmt.Println(m.Version)
	// hello.Hello()
}
