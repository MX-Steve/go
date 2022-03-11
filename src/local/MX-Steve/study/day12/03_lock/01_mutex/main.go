package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	var mutex sync.Mutex
	fmt.Println("locking (G0)")
	mutex.Lock()
	fmt.Println("locked (G0)")
	wg.Add(1)
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("locked (G%d)\n", i)
			time.Sleep(2 * time.Second)
			mutex.Unlock()
			fmt.Printf("unlocked (G%d)\n", i)
			wg.Done()
		}(i)
	}
	time.Sleep(7 * time.Second)
	fmt.Println("ready unlock (G0)")
	mutex.Unlock()
	fmt.Println("unlocked (G0)")
	wg.Wait()
}
