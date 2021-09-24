package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userLoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.html", nil)
}
func userRegisterPagePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register.html", nil)
}
func userLoginHandler(c *gin.Context) {
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
	c.Redirect(http.StatusMovedPermanently, "/v1/book/list")
}
func userRegisterHandler(c *gin.Context) {
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
}
