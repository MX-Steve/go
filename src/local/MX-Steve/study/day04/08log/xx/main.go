package main

import (
	"fmt"

	"local/MX-Steve/study/day04/08log/mylog"
)

// 自己的项目使用 mylog 这个包
func main() {
	f1 := mylog.NewFileLogger(mylog.CRATAL, "./", "test.log")
	// f1.Debug("这是一条测试的日志")
	fmt.Println("可以申请IPO")
	userID := 10
	f1.Cratal("id是 %d 的用户一直尝试登录", userID)
	// f1.Info("这是一条INFO级别的日志")
}
