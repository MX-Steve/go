package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(http.StatusOK, "book/book_list.tmpl", gin.H{
		"code": 0,
		"data": bookList,
	})
}
func newBookHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "book/new_book.html", nil)
}
func createBookHandler(c *gin.Context) {
	title := c.PostForm("title")
	price := c.PostForm("price")
	priceF, err := strconv.ParseFloat(price, 64)
	if err != nil {
		err = errors.New("invalid price parameter")
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	err = insertBook(title, priceF)
	if err != nil {
		err = errors.New("insert error, try again")
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func deleteBookHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		err = errors.New("invalid id parameter")
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	err = deleteBook(id)
	if err != nil {
		err = errors.New("delete book err,try again ")
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func modifyBookHandler(c *gin.Context) {
	idStr := c.PostForm("id")
	titleStr := c.PostForm("title")
	priceStr := c.PostForm("price")
	fmt.Println(idStr, titleStr, priceStr)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		msg := "invalid id parameter"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		msg := "invalid price parameter"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	err = modifyBook(id, titleStr, price)
	if err != nil {
		msg := "modify book err,try again "
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}
func modifyBookPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "book/modify_book.html", nil)
}
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
	c.Redirect(http.StatusMovedPermanently, "/book/list")
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
	c.Redirect(http.StatusMovedPermanently, "/user/login")
}
func main() {
	// init db
	err := initDB()
	if err != nil {
		panic(err)
	}
	// start gin
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/statics", "./statics")
	r.GET("/book/list", bookListHandler)
	r.GET("/book/new", newBookHandler)
	r.POST("/book/new", createBookHandler)
	r.GET("/book/delete", deleteBookHandler)
	r.GET("/book/modify", modifyBookPageHandler)
	r.POST("/book/modify", modifyBookHandler)
	r.GET("/user/login", userLoginPageHandler)
	r.GET("/user/register", userRegisterPagePageHandler)
	r.POST("/user/login", userLoginHandler)
	r.POST("/user/register", userRegisterHandler)
	r.Run(":9090")
}
