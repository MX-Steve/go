package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Gin with logrus
var log = logrus.New()

func initLogrus() (err error) {
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Error("fail to log to file")
		return
	}
	log.Out = file
	gin.SetMode(gin.ReleaseMode) // only use online
	// gin.DefaultWriter = log.Out
	gin.DefaultWriter = io.MultiWriter(log.Out, os.Stdout)
	log.Level = logrus.InfoLevel
	return
}
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "this is a test",
	})
}
func main() {
	err := initLogrus()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	log.Debug("This is a debug message")
	log.Info("trying to start server...")
	router.GET("/index", indexHandler)
	router.Run(":9090")
}
