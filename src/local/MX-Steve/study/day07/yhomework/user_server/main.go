package main

import (
	"local/MX-Steve/study/day07/yhomework/mylogger"
)

var logger mylogger.Logger

func main() {
	logger := mylogger.NewFileLogger("Debug", "./", "xxx.log")
	for {
		sb := "关大博"
		logger.Debug("%s是个日志Debug", sb)
		logger.Info("Info 这是一条测试数据")
		logger.Warn("Warn 这是一条测试数据")
		logger.Error("Error 这是一条测试数据")
	}
	// logger2 := mylogger.NewConsoleLogger("Info")
	// logger2.Debug("是个日志Debug")
	// logger2.Info("是个日志Info")
}
