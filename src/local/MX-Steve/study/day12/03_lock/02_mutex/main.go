package main

import (
	"fmt"
	"sync"
	"time"
)

type Book struct {
	BookName string
	L        *sync.Mutex
}

func (bk *Book) SetName(wg *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("unlock set name: ", name)
		bk.L.Unlock()
		wg.Done()
	}()
	bk.L.Lock()
	fmt.Println("lock set name: ", name)
	time.Sleep(time.Second * 1)
	bk.BookName = name
}

func main() {
	bk := Book{}
	bk.L = new(sync.Mutex)
	wg := &sync.WaitGroup{}
	books := []string{"《三国演义》", "《水浒传》", "《红楼梦》", "《西游记》"}
	for _, book := range books {
		wg.Add(1)
		go bk.SetName(wg, book)
	}
	wg.Wait()
}
