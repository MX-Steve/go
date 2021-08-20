package main 

import (
	"fmt"
)

func funcA(){
	fmt.Println("func a")
}

func funcB(){
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC(){
	fmt.Println("func c")
}

func main(){
	funcA()
	funcB()
	funcC()
}