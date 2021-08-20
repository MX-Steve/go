package main

import (
	"local/MX-Steve/study/day04/100/mylog"
)

func main() {
	f1 := mylog.NewFileLogger(mylog.INFO, ".", "test.log")
	f1.Debug("This is a test log file2: %d", 10)
	f1.Info("This is a test log file2: %d", 10)
	f1.Warn("This is a test log file2: %d", 10)
	f1.Error("This is a test log file2: %d", 10)
	f1.Cirtal("This is a test log file2: %d", 10)
}
