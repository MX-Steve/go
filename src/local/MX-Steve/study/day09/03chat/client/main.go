package main

import (
	"encoding/json"
	"fmt"
	"net"

	"local/MX-Steve/study/day09/03chat/server/protocol"
)

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}

func login(conn net.Conn) (err error) {
	var msg Message
	msg.Cmd = protocol.UserLogin
	var LoginCmd LoginCmd
	LoginCmd.Id = 1
	LoginCmd.Passwd = "123456789"
	data, err := json.Marshal(LoginCmd)
	if err != nil {
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}
	return
}
func main() {
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer conn.Close()
	err = login(conn)
	if err != nil {
		fmt.Println("login failed, err: ", err)
	}

	/*
		inputReader := bufio.NewReader(os.Stdin)
		for {
			input, _ := inputReader.ReadString('\n')
			trimmedInput := strings.Trim(input, "\n")
			if trimmedInput == "Q" {
				return
			}
			_, err = conn.Write([]byte(trimmedInput))
			if err != nil {
				return
			}
		}
	*/
}
