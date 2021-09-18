package main

import (
	"fmt"
	"io"
	"net"
)

// http client
func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			fmt.Println(string(buf[:n]))
			return
		}
		if err != nil {
			fmt.Println("receive data err,", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
