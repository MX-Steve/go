package main

import (
	"local/MX-Steve/study/day05/100/01logs/mylogger"
)

var logger mylogger.Logger

func main() {
	logger = mylogger.NewFileLogger("info", "./", "xxx.log")
	logger.Debug("%s is a good girl", "little pig pig")
	logger.Info("%s is a good girl", "little pig pig")
	logger.Warning("%s is a good girl", "little pig pig")
	logger.Error("%s is a good girl", "little pig pig")
	logger.Fatal("%s is a good girl", "little pig pig")
	logger = mylogger.NewConsoleLogger("info")
	logger.Debug("%s is a good girl", "little pig pig")
	logger.Info("%s is a good girl", "little pig pig")
	logger.Warning("%s is a good girl", "little pig pig")
	logger.Error("%s is a good girl", "little pig pig")
	logger.Fatal("%s is a good girl", "little pig pig")
}
