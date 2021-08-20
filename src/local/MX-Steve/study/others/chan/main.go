package main 

import (
	"fmt"
	"math/rand"
	"time"
)

func child(mychan chan int, index int){
	time.Sleep(time.Second*time.Duration(rand.Intn(10)))
	mychan <- index
}

func main(){
	var testchan = make(chan int,10)
	for i:=0;i<10;i++{
		go child(testchan,i)
	}
	for i:=0;i<10;i++{
		a := <- testchan
		fmt.Println("携程完成: ",a)
	}
	fmt.Println("主程完成")
}