package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "hello, how are you?"
		conn.Write([]byte(msg))
	}
}