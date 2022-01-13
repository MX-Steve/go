package main

import (
	"errors"
	"fmt"
	"net/http"
)

func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero")
		return
	}
	i = a / b
	return
}
func bar() (i int) {
	i = 0
	defer func() {
		fmt.Println(i)
	}()
	return 2
}
func do() error {
	res, err := http.Get("http://www.baidu.com")
	defer res.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	do()
	bar()
	foo(2, 0)
}
