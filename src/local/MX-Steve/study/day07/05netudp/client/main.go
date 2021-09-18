package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("connect server err,", err)
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte("do you have time?"))
	if err != nil {
		fmt.Println("send data err,", err)
		return
	}
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("response: ", string(buf[:n]))
}
