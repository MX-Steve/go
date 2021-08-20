package main

import "fmt"

type speaker interface {
	speak()
}
type mover interface {
	move()
}
type animal interface {
	speaker
	mover
}
type Cat struct {
	name string
}

func (c Cat) move() {
	fmt.Printf("%s can move\n", c.name)
}
func (c Cat) speak() {
	fmt.Printf("%s can speak\n", c.name)
}

func main() {
	var x animal
	fmt.Printf("%#v", x)
	x = Cat{name: "Tom cat"}
	x.move()
	x.speak()
	fmt.Println(x)
	fmt.Printf("%#v", x)
}
