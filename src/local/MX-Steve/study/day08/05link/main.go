package main

import "fmt"

// link operation

type student struct {
	name string
}

func (s student) learn() student {
	fmt.Printf("%s likes studying.\n", s.name)
	return s
}
func (s student) doHomework() student {
	fmt.Printf("%s likes doing homework.\n", s.name)
	return s
}
func main() {
	var stu1 = student{
		"San Zhang",
	}
	// stu1.learn()
	// stu1.doHomework()
	stu1.learn().doHomework()
}
