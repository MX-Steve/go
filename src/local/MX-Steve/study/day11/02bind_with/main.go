package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}
type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func someHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, "the body should be formA")
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, "the body should be formB")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "nothing",
		})
	}

}

func main() {
	router := gin.Default()
	router.Any("/some", someHandler)
	router.Run(":9090")
}
