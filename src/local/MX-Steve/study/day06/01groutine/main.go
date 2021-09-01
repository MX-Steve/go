package main

import (
	"fmt"
	"runtime"
	"sync"
)

// goroutine:
// 1. start routine about 2KB.
// 2. automatic capacity expansion when process is running.
// 3. don't need interactivate with kernel, runtime do that.

// GOMAXPROCS
// 1. use one cpu.(<go 1.5)
// 2. use all cpus.(>=go 1.5)
// 3. one thread can run many goroutines; many to many between thread and goroutine.

// start groutine
func hello() {
	defer wg.Done() // groutine complete; without this configure, it will cause dead lock.
	fmt.Println("hello,world")
}

// ex:2.a
func a() {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		fmt.Println("A:", i)
	}
}

//ex:2.b
func b() {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		fmt.Println("B:", i)
	}
}

// elegant waiting
var wg sync.WaitGroup

func main() {
	testIt := true
	if testIt {
		// EXAMPLE 2:
		wg.Add(2)
		runtime.GOMAXPROCS(1)
		// set the current process to use only one CPU.
		// m/n , set n=1.
		// m: goroutines's count; n: thread's count which really work
		go a()
		go b()
		wg.Wait()
	} else {
		fmt.Println("finished...")
		// EXAMPLE 1:
		defer fmt.Println("finish main func")
		for i := 0; i < 10; i++ {
			wg.Add(1)  // count groutine
			go hello() // 1. start a new groutine; 2. run hello function in this new groutine
		}
		fmt.Println("hello main func")
		// time.Sleep(time.Second)
		// wait to the new groutine
		wg.Wait() // wait for all groutines to complete
	}

}
