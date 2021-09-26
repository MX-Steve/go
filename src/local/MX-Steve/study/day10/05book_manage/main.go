package main

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

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
	r.GET("/", userLoginHandler)
	v1 := r.Group("/v1")
	{
		v1User := v1.Group("/user")
		{
			v1User.Any("/login", userLoginHandler)
			v1User.Any("/register", userRegisterHandler)
		}
		v1Book := v1.Group("/book")
		template.ParseFiles("templates/book/book_list.html", "templates/book/head.html", "templates/book/head.html",
			"templates/book/foot.html")
		template.ParseFiles("templates/book/modify_book.html", "templates/book/head.html", "templates/book/head.html",
			"templates/book/foot.html")
		template.ParseFiles("templates/book/new_book.html", "templates/book/head.html", "templates/book/head.html",
			"templates/book/foot.html")
		{
			v1Book.GET("/list", bookListHandler)
			v1Book.Any("/new", createBookHandler)
			v1Book.GET("/delete", deleteBookHandler)
			v1Book.Any("/modify", modifyBookHandler)
		}
	}
	v2 := r.Group("/v2")
	{
		v2User := v2.Group("/user")
		{
			v2User.GET("/login", userLoginPageV2Handler)
		}
	}
	r.Run(":9090")
}
