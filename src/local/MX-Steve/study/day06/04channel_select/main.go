package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel select: Multiplexing

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
var doneChan chan struct{}

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
	for tmp := range ch {
		sum := calc(tmp.num)
		retObj := &result{
			item: *tmp,
			sum:  sum,
		}
		resultChain <- retObj
	}
	doneChan <- struct{}{}
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

func closeResult(ch chan struct{}, resultChain chan *result) {
	for i := 0; i < 20; i++ {
		<-ch
	}
	close(ch)
	close(resultChan)
}

// printResult
func printResult(resultChain chan *result) {
	for ret := range resultChain {
		fmt.Printf("id:%v, num:%v, sum:%v\n", ret.item.id, ret.item.num, ret.sum)
		time.Sleep(time.Second)
	}

}

func startWorker(n int, ch chan *item, resultChain chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resultChan)
	}
}

var ch = make(chan string, 100)
var ch2 = make(chan string, 100)

func f1(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("f1:%d", i)
		time.Sleep(time.Millisecond * 50)
	}
}
func f2(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("f2:%d", i)
		time.Sleep(time.Millisecond * 100)
	}
}

var ch30 = make(chan int, 1)

func ex3(ch30 chan int) {

	for i := 0; i < 10; i++ {
		select {
		case ch30 <- i:
		case ret := <-ch30:
			fmt.Println(ret)
		}
	}
}

func main() {
	testIt := true
	if testIt {
		fmt.Println("starting...")
		// EX1
		itemChan = make(chan *item, 100)
		resultChan = make(chan *result, 100)
		doneChan = make(chan struct{}, 20)
		go producer(itemChan)
		// go consumer(itemChan, resultChan)
		startWorker(20, itemChan, resultChan)
		go closeResult(doneChan, resultChan)
		printResult(resultChan)
		// EX3
		ex3(ch30)
	} else {

		// EX2
		go f1(ch)
		go f2(ch2)
		for {
			select {
			case <-ch:
				fmt.Println(<-ch)
			case <-ch2:
				fmt.Println(<-ch2)
			default:
				fmt.Println("no data now")
				time.Sleep(time.Millisecond * 50)

			}
		}

	}

}
