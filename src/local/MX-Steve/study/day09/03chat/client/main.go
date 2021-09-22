package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"local/MX-Steve/study/day09/03chat/server/protocol"
)

func login(conn net.Conn) (err error) {
	var msg protocol.Message
	msg.Cmd = protocol.UserLogin
	var LoginCmd protocol.LoginCmd
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
	var buf [4]byte
	packLen := uint32(len(data))
	// buffer := bytes.NewBuffer(buf[:])
	// err = binary.Write(buffer, binary.BigEndian, packLen)
	binary.BigEndian.PutUint32(buf[0:4], packLen)
	// if err != nil {
	// 	fmt.Println("write package len failed")
	// 	return
	// }
	n, err := conn.Write(buf[:])
	// n, err := conn.Write(buffer.Bytes())
	if err != nil || n != 4 {
		fmt.Println("write data failed")
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("write data failed")
		return
	}

	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}
	// var buf [8192]byte
	// conn.Read(buf[:])
	msg, err = readPackage(conn)
	if err != nil {
		fmt.Println("read package failed, err ", err)
	}
	fmt.Println(msg)
	return
}

func readPackage(conn net.Conn) (msg protocol.Message, err error) {
	var buf [8192]byte
	n, _ := conn.Read(buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}

	// buffer := bytes.NewBuffer(p.buf[0:8])
	// var packLen uint32
	// err = binary.Read(buffer, binary.BigEndian, &packLen)
	packLen := binary.BigEndian.Uint32(buf[0:4])
	if err != nil {
		fmt.Println("read package len failed")
		return
	}
	n, _ = conn.Read(buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read header failed")
		return
	}
	err = json.Unmarshal(buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err: ", err)
	}
	return
}
func main() {
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = login(conn)
	if err != nil {
		fmt.Println("login failed, err: ", err)
	}

}
