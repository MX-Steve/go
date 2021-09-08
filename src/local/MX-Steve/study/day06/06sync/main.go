package main

import (
	"fmt"
	"sync"
	"time"
)

// concurrent and lock

var x int64
var wg sync.WaitGroup

// define a Mutual exclusion lock
var lock sync.Mutex

// Concurrent and lock
func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // add lock
		x = x + 1
		lock.Unlock() // remove lock
	}
	wg.Done()
}

var y int64
var rwLock sync.RWMutex

// Use RWMutex
// 1. Concurrent read, add write lock when write values.
// 2. the count for reading more than the count for writing

func read() {
	defer wg.Done()
	rwLock.RLock()
	// lock.Lock()
	// fmt.Println(y)
	time.Sleep(time.Millisecond * 1)
	rwLock.RUnlock()
	// lock.Unlock()

}
func write() {
	defer wg.Done()
	rwLock.Lock()
	// lock.Lock()
	y = y + 1
	time.Sleep(time.Millisecond * 5)
	rwLock.Unlock()
	// lock.Unlock()
}

// Mutual exclusion lock about read and write
func main() {
	testIt := true
	if testIt {
		fmt.Println("starting.")
		start := time.Now()
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go write()
		}
		for i := 0; i < 10000; i++ {
			wg.Add(1)
			go read()
		}
		wg.Wait()
		end := time.Now()
		fmt.Printf("spend %v.\n", end.Sub(start))
	} else {
		// EX1
		wg.Add(2)
		go add()
		go add()
		wg.Wait()
		fmt.Println(x)
	}

}
