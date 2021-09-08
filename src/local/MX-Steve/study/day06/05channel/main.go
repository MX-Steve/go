package main

import (
	"fmt"
)

// Concurrent and Lock
// can send and receive values
// ch chan<- string : only send values
// ch <-chan string : only receive values
func getInput(ch chan string) {
	var input string
	fmt.Scan(&input)
	ch <- input
}
func main() {
	exitChan := make(chan string, 1)
	getInput(exitChan)
	fmt.Println("in main")
	fmt.Println(<-exitChan)
}
