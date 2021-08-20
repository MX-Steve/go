package main

import "fmt"

var today = "sunday"

const Week = 7

type Student struct {
	Name string
}

func init() {
	fmt.Println("init in main")
}

func main() {
	fmt.Println(today)
}
