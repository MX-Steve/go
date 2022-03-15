package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "local/MX-Steve/study/day14/06_marshal/05_buffer/server/djentry"
)

func main() {
	usermsg := &pb.UserInfo{
		UserType: 1,
		UserName: "Jok",
		UserInfo: "I am a worker!",
	}
	userdata, err := proto.Marshal(usermsg)
	if err != nil {
		fmt.Println("Marshaling error: ", err)
	}
	encodingmsg := &pb.UserInfo{}
	err = proto.Unmarshal(userdata, encodingmsg)
	if err != nil {
		fmt.Println("Unmarshaling error: ", err)
	}
	fmt.Printf("GetUserType: %d\n", encodingmsg.GetUserType())
	fmt.Printf("GetUserName: %s\n", encodingmsg.GetUserName())
	fmt.Printf("GetUserInfo: %s\n", encodingmsg.GetUserInfo())
}
