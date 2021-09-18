package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// tcp client
func main() {
	// 1. find server by address and port
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("connect server err,", err)
		return
	}
	defer conn.Close()
	// fmt.Fprintln(conn, "do you have time?")
	// var input string
	// fmt.Scanln(&input)
	rd := bufio.NewReader(os.Stdin)
	str, err := rd.ReadString('\n')
	if err != nil {
		fmt.Println("occur err,", err)
	}
	_, err = conn.Write([]byte(str))
	if err != nil {
		fmt.Println("send data err,", err)
		return
	}
}
