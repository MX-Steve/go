package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var onlyOnce sync.Once

func t1() {
	var m = make(map[string]int)
	fmt.Println(m)
	m["score"] = 89
	fmt.Println(m)
}

func f1(a int) {
	fmt.Println(a)
}
func closure(x int) func() {
	return func() {
		f1(x)
	}
}

var smap = sync.Map{}

func t() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}

}
func main() {
	f2 := func() bool {
		return false
	}
	switch f2() {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
	t()
	wg := sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			smap.Store(key, n)
			value, _ := smap.Load(key)
			fmt.Printf("k:%v,v:%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

	f := closure(10)
	onlyOnce.Do(f)

	t1()
}
