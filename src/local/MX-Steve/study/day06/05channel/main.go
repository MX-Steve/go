package main

import (
	"fmt"
	"os"
)

func getInput() {
	tmp := [1]byte{}
	n, _ := os.Stdin.Read(tmp[:])
	fmt.Println(n)
	fmt.Println(123)
}
func main() {
	getInput()
}
