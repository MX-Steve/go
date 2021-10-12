package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func initLogrus() (err error) {
	log.Formatter = &logrus.JSONFormatter{}
	file, err := os.OpenFile("./bms.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Error("fail to log to file")
		return
	}
	log.Out = file
	log.Level = logrus.InfoLevel
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	return
}

func cookieMiddleware(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		toPath := fmt.Sprintf("%s?next=%s", "/v1/user/login", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, toPath)
		// c.HTML(http.StatusOK, "user/login.html", gin.H{
		// 	"msg": "you have no cookie username or expired",
		// })
		return
	}
	c.Set("username", username)
	c.Next()
}

func main() {
	// init db
	err := initDB()
	if err != nil {
		panic(err)
	}
	err = initLogrus()
	if err != nil {
		panic(err)
	}
	log.Info("trying to start server...")
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
		v1Book := v1.Group("/book", cookieMiddleware)
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
		v1Publisher := v1.Group("/publisher", cookieMiddleware)
		template.ParseFiles("templates/publisher/publisher_list.html", "templates/book/head.html", "templates/book/head.html",
			"templates/book/foot.html")
		template.ParseFiles("templates/publisher/modify_publisher.html", "templates/book/head.html", "templates/book/head.html",
			"templates/book/foot.html")
		template.ParseFiles("templates/publisher/new_publisher.html", "templates/book/head.html", "templates/book/head.html",
			"templates/book/foot.html")
		{
			v1Publisher.GET("/list", publisherListHandler)
			v1Publisher.Any("/new", createPublisherHandler)
			v1Publisher.GET("/delete", deletePublisherHandler)
			v1Publisher.Any("/modify", modifyPublisherHandler)
		}
		v1Word := v1.Group("/word")
		{
			v1Word.Any("/list", wordListHandler)
			v1Word.Any("/new", createWordHandler)
			v1Word.GET("/delete", deleteWordHandler)
			v1Word.Any("/modify", modifyWordHandler)
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
