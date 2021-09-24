package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// html render
func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "log website",
	})
}
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "index website",
	})
}
func main() {
	router := gin.Default()
	// htmls
	router.LoadHTMLGlob("templates/*")
	// statics
	router.Static("/st", "./statics")
	router.GET("/login", loginHandler)
	router.GET("/index", indexHandler)
	router.Run(":9090")
}
