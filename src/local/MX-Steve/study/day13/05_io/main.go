package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	// b := make([]byte, 1024)
	// f, err := os.Open("./tt.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = f.Read(b)
	// f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(b))
	file := "./go.mod"
	start := time.Now()
	read1(file)
	t1 := time.Now()
	fmt.Printf("Cost time: %v\n", t1.Sub(start))
	read2(file)
	t2 := time.Now()
	fmt.Printf("Cost time: %v\n", t2.Sub(t1))
	read3(file)
	t3 := time.Now()
	fmt.Printf("Cost time: %v\n", t3.Sub(t2))
}

func read1(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
}

func read2(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
}

func read3(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	_, err = ioutil.ReadAll(fi)
	if err != nil {
		fmt.Println(err)
	}
}
