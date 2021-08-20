package main

import "fmt"

type animal interface {
	speak()
	move()
}

// 1. value receiver
type cat struct {
	name string
}

func (c cat) speak() {
	fmt.Println("miaomiao")
}
func (c cat) move() {
	fmt.Println("can move")
}

// 2. pointer receiver
type dog struct {
	name string
}

func (c *dog) speak() {
	fmt.Println("wangwang")
}
func (c *dog) move() {
	fmt.Println("can move")
}

func main() {
	testIt := true
	if testIt {
		fmt.Println("Start...")

	} else {
		fmt.Println("Finished...")
		// 1. value receiver
		fmt.Println("Start...")
		var x animal
		hh := cat{"huahua"}
		x = hh
		tom := &cat{"jialaoliehefentaotao"}
		fmt.Printf("%T\n", tom)
		fmt.Println(tom)
		tom.speak()
		x = tom
		fmt.Println(x)
		// 2. pointer receiver
		var y animal
		dd := &dog{"Lucy"}
		y = dd
		fmt.Println(y)
	}

}
