package main

import (
	"bufio"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	rd := bufio.NewReader(os.Stdin)
	str, err := rd.ReadString('\n')
	if err != nil {
		log.Info(err)
	}
	_, err = conn.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
	}
}
