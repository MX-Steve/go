package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// cookie demo
var log = logrus.New()

func initLogrus() (err error) {
	log.Formatter = &logrus.JSONFormatter{}
	log.Level = logrus.InfoLevel
	// file, err := os.OpenFile("cookie.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	log.Out = os.Stdout
	gin.SetMode(gin.ReleaseMode)
	// gin.DefaultWriter = io.MultiWriter(log.Out, os.Stdout)
	gin.DefaultWriter = log.Out
	return
}
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		toPath := c.DefaultQuery("next", "/index")
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "username or password can't be null",
			})
			return
		}
		if u.Username == "lisi" && u.Password == "123456" {
			c.SetCookie("username", u.Username, 20, "/", "127.0.0.1", false, true)
			c.Redirect(http.StatusMovedPermanently, toPath)
			return
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "username or password error",
			})
			return
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}

}

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func homeHandler(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		// c.Redirect(http.StatusMovedPermanently, "/login")
		// return
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err": "you have no cookie username or expired",
		})
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
	})
}
func vipHandler(c *gin.Context) {
	tmpUsername, ok := c.Get("username")
	if !ok {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err": "you have no cookie username or expired",
		})
		return
	}
	username, ok := tmpUsername.(string)
	if !ok {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err": "you have no cookie username or expired",
		})
		return
	}
	c.HTML(http.StatusOK, "vip.html", gin.H{
		"msg":      "welcome to vip",
		"username": username,
	})
}
func cookieMiddleware(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		// c.Redirect(http.StatusMovedPermanently, "/login")
		// return

		toPath := fmt.Sprintf("%s?next=%s", "/login", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, toPath)
		return
	}
	c.Set("username", username)
	c.Next() // continue next functions
}
func main() {
	err := initLogrus()
	if err != nil {
		panic(err)
	}
	log.Info("trying to start server")
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	router.Any("/login", loginHandler)
	router.GET("/home", homeHandler)
	router.GET("/index", indexHandler)
	router.GET("/vip", cookieMiddleware, vipHandler)
	router.Run(":9090")
}
