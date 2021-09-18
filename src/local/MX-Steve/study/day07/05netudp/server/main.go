package main

import (
	"fmt"
	"net"
)

func process(listenner *net.UDPConn) {
	defer listenner.Close()
	for {
		var buf [1024]byte
		n, addr, err := listenner.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("receive data err,", err)
		}
		fmt.Printf("receive data from %v: %v\n", addr, string(buf[:n]))
		// send data
		_, err = listenner.WriteToUDP([]byte("go out."), addr)
		if err != nil {
			fmt.Println("response err,", err)
			return
		}
	}
}

// udp server
func main() {
	listenner, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("start server err, ", err)
		return
	}
	process(listenner)
}
