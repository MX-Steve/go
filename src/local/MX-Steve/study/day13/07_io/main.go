package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	buf := bufio.NewReaderSize(sr, 0)
	b := make([]byte, 10)
	fmt.Println("==", buf.Buffered())
	S, _ := buf.Peek(5)
	fmt.Printf("%d == %q\n", buf.Buffered(), S)
	nn, er := buf.Discard(3)
	fmt.Println(nn, er)
	for n, err := 0, error(nil); err == nil; {
		fmt.Printf("Buffered:%d == Size:%d==n:%d== b[:n] %q == err:%v\n", buf.Buffered(), buf.Size(), n, b[:n], err)
		n, err = buf.Read(b)
		fmt.Printf("Buffered:%d ==Size:%d== n:%d==  b[:n] %q ==  err: %v == s: %s\n", buf.Buffered(), buf.Size(), n, b[:n], err, S)
	}
	fmt.Printf("%d ==  %q\n", buf.Buffered(), S)
}
