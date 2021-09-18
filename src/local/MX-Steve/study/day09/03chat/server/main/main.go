package main

import "time"

/*
0. porpose
	1. chatting platform which many people can chat
	2. chat point by point
	3. user login & register
一. server
1. user managerment
	uuid: int
	passwd: string
	name: string
	gender: string
	profile: url
	onlineTime: string
	online: bool
	data storage: redis hash: users
2. user operation
	1. send message
	2. receive message
3. user register & login
4. storage data where user is offline
二. client
1. user register
2. user login
3. send message
4. get user list
三.communication protocol
1. [0:4] length
2. [] json
*/
func main() {
	initRedis("localhost:6379", 16, 1024, time.Second*300)
	initUserMgr()
	runServer("0.0.0.0:10000")
}
