package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// producer and consumer model
// use goroutine and channel
// producer: generate random
// consumer: calculate the sum of each digit
// 1 producer, 20 consumer

type item struct {
	id  int64
	num int64
}
type result struct {
	item
	sum int64
}

var itemChan chan *item
var resultChan chan *result

// producer
func producer(ch chan *item) {
	var id int64
	for i := 0; i < 10; i++ {
		id++
		ret := rand.Int63()
		tmp := &item{
			id:  id,
			num: ret,
		}
		ch <- tmp
	}
	close(ch)
}

// var wg sync.WaitGroup

// consumer
func consumer(ch chan *item, resultChain chan *result) {
	defer wg.Done()
	for tmp := range ch {
		sum := calc(tmp.num)
		retObj := &result{
			item: *tmp,
			sum:  sum,
		}
		resultChain <- retObj
	}
}

// calc: get sum
func calc(num int64) int64 {
	var sum int64
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

// printResult
func printResult(resultChain chan *result) {
	for ret := range resultChain {
		fmt.Printf("id:%v, num:%v, sum:%v\n", ret.item.id, ret.item.num, ret.sum)
		time.Sleep(time.Second)
	}
	// for {
	// 	ret := <-resultChain
	// 	if ret == nil {
	// 		break
	// 	}
	// 	fmt.Printf("id:%v, num:%v, sum:%v\n", ret.item.id, ret.item.num, ret.sum)
	// }
}

func startWorker(n int, ch chan *item, resultChain chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resultChan)
	}
}

var wg sync.WaitGroup

func main() {
	itemChan = make(chan *item, 100)
	resultChan = make(chan *result, 100)
	go producer(itemChan)
	// go consumer(itemChan, resultChan)
	wg.Add(20)
	startWorker(20, itemChan, resultChan)
	go printResult(resultChan)
	wg.Wait()
	close(resultChan)

}
