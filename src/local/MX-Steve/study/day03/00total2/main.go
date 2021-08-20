package main

import "fmt"

// method
type People struct {
	name   string
	gender string
}

func (p *People) dream() {
	fmt.Printf("%s的梦想是%s\n", p.name, p.gender)
}

type MyInt int

// MyInt.sayHi
func (m *MyInt) sayHi() {
	fmt.Println("hi...")
}

// anonymouse struct
type Student struct {
	name string
	string
	int
}

// Nesting of structures
type Address struct {
	province string
	city     string
}
type Teacher struct {
	name string
	age  int
	addr Address
}

func main() {
	// var p = &People{"lisi", "basketball"}
	// p.dream()
	// var a MyInt
	// fmt.Println(a)
	// a.sayHi()
	// var stu1 = Student{
	// 	name: "lisi",
	// }
	// fmt.Println(stu1)
	// fmt.Println(stu1.name)
	// fmt.Println(stu1.string)
	// fmt.Println(stu1.int)
	var t1 = Teacher{
		name: "lisi",
		age:  28,
		addr: Address{
			province: "beijing",
			city:     "anh",
		},
	}
	fmt.Println(t1)
}
