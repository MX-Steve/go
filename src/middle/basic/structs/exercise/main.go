package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	ce[1].age = 999
}

func main() {
	var ce []student
	ce = []student{
		{1, "xiaoming", 22},
		{2, "xiaohong", 23},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
	c1 := make(map[int]student)
	c1[1] = student{1, "xiaoli", 22}
	c1[2] = student{2, "hong", 22}
	fmt.Println(c1)
	delete(c1, 2)
	fmt.Println(c1)
}
