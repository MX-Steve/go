package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func shopIndexHandler(c *gin.Context) {
	time.Sleep(time.Second * 1)
	fmt.Println(c.MustGet("key").(string))
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/index",
	})
}
func shopHomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/home",
	})
}
func castTime(c *gin.Context) {
	start := time.Now()
	c.Set("key", "value")
	c.Next()
	cast := time.Since(start)
	fmt.Println("cast: ", cast)
}

func main() {
	r := gin.Default()
	// middleware
	r.Use(castTime)
	// shoppingGroup := r.Group("/shopping", castTime)
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("index", shopIndexHandler)
		shoppingGroup.GET("home", shopHomeHandler)
	}
	r.Run(":9090")
}
