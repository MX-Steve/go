package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func onces() {
	fmt.Println("ones")
}
func onced() {
	fmt.Println("onced")
}

func main() {
	// demo 2
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("v: ", v, "----i:", i)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			once.Do(onced)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(4000)
	// demo 1
	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}(i)
	// }
	// wg.Wait()
	// fmt.Println("程序结束")

}
