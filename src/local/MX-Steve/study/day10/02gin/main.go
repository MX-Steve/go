package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
separate between frontend and backend
MVC:
	modle: data
	view: show
	controller: operation
MVVM:
	1. MV:
	2. VM:
	3. steps:
		1. use url and method to run function
		2. get data from db
		3. serialize data into json format
		4. send json to frontend
*/
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "this is a index page",
	})
}
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world!",
		})
	})
	r.GET("/index", indexHandler)
	r.Run(":9090")
}
