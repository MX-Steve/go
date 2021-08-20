package main

import (
	"fmt"
)

func main() {
	age := 19
	if age > 18 {
		fmt.Println("赌场开业啦")
	} else if age < 18 {
		fmt.Println("warning")
	} else {
		fmt.Println("you are adult")
	}
}
