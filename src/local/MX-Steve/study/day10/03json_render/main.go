package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "This is index page",
	})
}
func helloHandler(c *gin.Context) {
	type userInfo struct {
		Name     string
		Password string `json:"pwd"`
	}
	u1 := userInfo{
		Name:     "lisi",
		Password: "123456",
	}
	c.JSON(http.StatusOK, u1)
}
func xmlHandler(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"msg": "xml",
	})
}
func yamlHandler(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{
		"message": "ok",
		"status":  http.StatusOK,
	})
}
func main() {
	r := gin.Default()
	r.GET("/index", indexHandler)
	r.GET("/hello", helloHandler)
	r.GET("/xml", xmlHandler)
	r.GET("/yaml", yamlHandler)
	r.Run(":9090")
}
