package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func wordListHandler(c *gin.Context) {
	wordList, err := queryAllWords()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(http.StatusOK, "word/word_list.tmpl", gin.H{
		"code": 0,
		"data": wordList,
	})
}

func createWordHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		ename := c.PostForm("ename")
		cname := c.PostForm("cname")
		splits := c.PostForm("splits")
		memory := c.PostForm("memory")
		example := c.PostForm("example")
		err := insertWord(ename, cname, splits, memory, example)
		if err != nil {
			err = errors.New("insert error, try again")
			c.JSON(http.StatusOK, gin.H{
				"msg": err,
			})
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/v1/word/list")
	} else {
		c.HTML(http.StatusOK, "word/new_word.html", nil)
	}

}

func deleteWordHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		err = errors.New("invalid id parameter")
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	err = deleteWord(id)
	if err != nil {
		err = errors.New("delete book err,try again ")
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/v1/word/list")
}

func modifyWordHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		idStr := c.PostForm("id")
		ename := c.PostForm("ename")
		cname := c.PostForm("cname")
		splits := c.PostForm("splits")
		memory := c.PostForm("memory")
		example := c.PostForm("example")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			msg := "invalid id parameter"
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
			return
		}
		err = modifyWord(id, ename, cname, splits, memory, example)
		if err != nil {
			msg := "modify word err,try again "
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/v1/word/list")
	} else {
		idStr := c.Query("id")
		if len(idStr) == 0 {
			c.String(http.StatusBadRequest, "invalid request")
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid request")
			return
		}
		wordObj, err := getWordById(id)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid word id")
			return
		}
		c.HTML(http.StatusOK, "word/modify_word.html", wordObj)
	}
}
