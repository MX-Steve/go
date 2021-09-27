package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func logEx1() {
	var log1 = log.New()
	file, err := os.OpenFile("./mes.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Warn("fail to log to file")
		return
	}
	log1.Out = file
	log1.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	log1.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
	log.WithFields(log.Fields{
		"name": "steve",
		"age":  28,
	}).Warn("This is a warning message.")
	log.Error("This is a error message.")

}
func logEx2() {
	requestLevel := log.WithFields(log.Fields{
		"project": "platform",
		"author":  "steve",
	})
	requestLevel.Info("This is a info message")
	log.SetFormatter(&log.JSONFormatter{})
	requestLevel.Warn("This is a warn message")
}
func main() {
	testIt := true
	if testIt {
		logEx2()
	} else {
		logEx1()
	}
}
