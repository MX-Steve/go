package main

import (
	"fmt"
	"time"
)

var resChan = make(chan int)

func test() {
	select {
	case data := <-resChan:
		doData(data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}
}

func doData(data int) {
	fmt.Println(data)
}

func main() {
	test()
}
