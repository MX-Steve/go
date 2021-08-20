package main

import "fmt"

func main() {
	// const (
	// 	school    string = "beijing"
	// 	className string = "class2"
	// 	age       int    = 222
	// )
	// fmt.Println(school, className, age)

	const (
		aa = iota
		bb
		cc
		dd
		_
		ee
	)
	fmt.Println(aa, bb, cc, dd, ee)
}
