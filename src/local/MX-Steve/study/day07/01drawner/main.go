package main

import "fmt"

func f1(ch chan int) {
	<-ch
}
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 1)
	select {
	case ch1 <- 100:
		fmt.Println("1111")
	case <-ch2:
		fmt.Println("<-ch2")
	default:
		fmt.Println("default")
	}

	ch := make(chan int)
	go f1(ch)
	ch <- 10
	fmt.Println("in main")
}
