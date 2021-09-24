package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context) {
	fileObj, err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	fmt.Printf("文件名: %s\n", fileObj.Filename)
	filePath := fmt.Sprintf("./%s", fileObj.Filename)
	c.SaveUploadedFile(fileObj, filePath)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./upload.html")
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	liveGroup := r.Group("/live")
	{
		liveGroup.GET("/index", liveIndexHandler)
		liveGroup.GET("/home", liveHomeHandler)
	}
	v1 := r.Group("/v1")
	{
		v1Shopping := v1.Group("/shopping")
		{
			v1Shopping.GET("/index", shopIndexHandler)
			v1Shopping.GET("/upload", func(c *gin.Context) {
				c.HTML(http.StatusOK, "upload.html", nil)
			})
			v1Shopping.POST("/upload", uploadHandler)
		}
	}
	r.Run(":9090")
}
