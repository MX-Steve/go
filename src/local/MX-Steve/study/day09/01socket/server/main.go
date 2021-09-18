package main

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("start server...")
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Info(err)
			continue
		}
		go process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			log.Info(err)
			return
		}
		// log.Info("read: ", string(buf))
		log.Info(strings.TrimSpace(string(buf)))
	}
}
