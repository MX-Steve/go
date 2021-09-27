package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func loginHander(c *gin.Context) {
	if c.Request.Method == "POST" {
		// operate data
		username := c.PostForm("username")
		password := c.PostForm("password")
		err := Login(username, password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "username or password err.",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "login ok",
		})
	} else {
		// login.html
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func login2Handler(c *gin.Context) {
	if c.Request.Method == "POST" {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusOK, err.Error())
			return
		}
		err = Login(u.Username, u.Password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "username or password err.",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "login ok",
		})
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func searchHandler(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func postHandler(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	day := c.Param("day")
	c.JSON(http.StatusOK, gin.H{
		"year":  year,
		"month": month,
		"day":   day,
	})
}

type UserInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func bookDetailHandler(c *gin.Context) {
	tmpBookId := c.Param("id")
	if len(tmpBookId) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err": "invalid param",
		})
		return
	}
	bookId, err := strconv.ParseInt(tmpBookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "invalid param",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": bookId,
	})
}
func main() {
	router := gin.Default()
	err := initDb()
	if err != nil {
		panic(err)
	}
	router.LoadHTMLGlob("./templates/*")
	router.Static("/statics", "./statics")
	router.GET("/", indexHandler)
	v1 := router.Group("/v1")
	{
		v1.GET("/index", indexHandler)
	}
	router.Any("/login", loginHander)
	router.Any("/login2", login2Handler)
	router.GET("/posts/:year/:month/:day", postHandler)
	router.GET("/search", searchHandler)
	router.GET("/book/:id", bookDetailHandler)
	router.Run(":9090")

}
