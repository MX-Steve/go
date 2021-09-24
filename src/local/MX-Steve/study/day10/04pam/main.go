package main

// parameters
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func queryStringHandler(c *gin.Context) {
	name := c.Query("name")
	password := c.DefaultQuery("password", "123456")
	c.JSON(http.StatusOK, gin.H{
		"name":     name,
		"password": password,
	})
}
func formHandler(c *gin.Context) {
	// get data from form
	name := c.DefaultPostForm("name", "root")
	passd := c.DefaultPostForm("password", "ssww")
	c.JSON(http.StatusOK, gin.H{
		"name":     name,
		"password": passd,
	})
}
func paramsHandler(c *gin.Context) {
	// get url params
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"action": action,
	})
}
func main() {
	r := gin.Default()
	r.GET("/query_string", queryStringHandler)
	r.POST("/form", formHandler)
	r.GET("/book/:action", paramsHandler)
	r.Run(":9090")
}
