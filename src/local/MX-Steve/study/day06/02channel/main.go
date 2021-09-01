package main

import "fmt"

// ex2:1
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// ex3:1
func recv(ch chan bool) {
	ret := <-ch
	fmt.Println(ret)
}

// ex4:1
func fab(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// ex5:1
func fib2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// ex:6:1
func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // don't must do that
}

// channel: a typed conduit through which you can send and receive values with channel operator, <- .
func main() {
	testIt := true
	if testIt {
		// EXAMPLE 6
		ch60 := make(chan int, 10)
		go send(ch60)
		for {
			ret, ok := <-ch60 // when no values and chan was closed.
			if !ok {
				break
			}
			fmt.Println(ret)
		}
	} else {
		fmt.Println("finished...")
		// EXAMPLE 1
		// define a type of chan which send and receive int values
		var ch1 chan int
		// chan is a type of reference
		fmt.Println(ch1)
		// initialize chan with make
		ch2 := make(chan int, 2)
		// fmt.Println(ch1)
		// method: send;receive,close
		// send and receive: <-
		ch2 <- 10
		ch2 <- 20
		ret := <-ch2
		fmt.Println(ret)
		fmt.Println(<-ch2)
		close(ch2)
		// you can also receive values 0 (int) after closing channel, but can't send values.

		// EXAMPLE 2
		s := []int{2, 3, 1, 55, 62, 34, 78}
		c := make(chan int)
		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:], c)
		x, y := <-c, <-c
		fmt.Println(x, y, x+y)

		// EXAMPLE 3
		// has no buffer channel
		ch30 := make(chan bool)
		go recv(ch30)
		ch30 <- true
		// has buffer channel
		ch31 := make(chan int, 2)
		ch31 <- 10
		fmt.Println(<-ch31)

		// EXAMPLE 4
		c40 := make(chan int, 10)
		go fab(cap(c40), c40)
		for i := range c40 {
			fmt.Println(i)
		}

		// EXAMPLE 5
		c50 := make(chan int)
		quit := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(<-c50)
			}
			quit <- 0
		}()
		fib2(c50, quit)
	}
}
