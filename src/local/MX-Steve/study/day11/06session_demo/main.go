package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func initLogrus() (err error) {
	log.Formatter = &logrus.JSONFormatter{}
	log.Level = logrus.InfoLevel
	file, err := os.OpenFile("session.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.Out = file
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	return
}
func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "lisi" && password == "123456" {
			c.Redirect(http.StatusMovedPermanently, "/index")
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": "username or password err",
		})
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func main() {
	err := initLogrus()
	if err != nil {
		panic(err)
	}
	log.Info("try to start server...")
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	router.Any("/login", loginHandler)
	router.Any("/index", indexHandler)
	router.Run(":9090")
}
