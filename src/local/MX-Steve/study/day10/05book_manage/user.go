package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userLoginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		toPath := c.DefaultQuery("next", "/v1/book/list")
		username := c.PostForm("username")
		password := c.PostForm("password")
		err := userLogin(username, password)
		if err != nil {
			msg := "用户名或密码错误"
			c.HTML(http.StatusOK, "user/login.html", gin.H{
				"msg": msg,
			})
			return
		}
		c.SetCookie("username", username, 3600*24, "/", "127.0.0.1", false, true)
		c.Redirect(http.StatusMovedPermanently, toPath)
	} else {
		c.HTML(http.StatusOK, "user/login.html", nil)
	}

}
func userRegisterHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		err := userRegister(username, password)
		if err != nil {
			msg := "用户名已存在"
			c.HTML(http.StatusOK, "user/register.html", gin.H{
				"msg": msg,
			})
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/v1/user/login")
	} else {
		c.HTML(http.StatusOK, "user/register.html", nil)
	}

}
