package main

import (
	"io"
	"net"

	log "github.com/sirupsen/logrus"
)

// lookup baidu website
func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\r\n"
	msg += "Host: www.baidu.com\r\n"
	msg += "Connection: keep-alived\r\n"
	msg += "\r\n\r\n"
	_, err = io.WriteString(conn, msg)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			break
		}
		log.Info(string(buf[0:count]))
	}
}
