package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func publisherListHandler(c *gin.Context) {
	publisherList, err := queryAllPublisher()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(http.StatusOK, "publisher/publisher_list.tmpl", gin.H{
		"code": 0,
		"data": publisherList,
	})
}

func createPublisherHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		province := c.PostForm("province")
		city := c.PostForm("city")
		name := c.PostForm("name")

		err := insertPublisher(province, city, name)
		if err != nil {
			err = errors.New("insert error, try again")
			c.JSON(http.StatusOK, gin.H{
				"msg": err,
			})
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/v1/publisher/list")
	} else {
		c.HTML(http.StatusOK, "publisher/new_publisher.html", nil)
	}

}

func deletePublisherHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		err = errors.New("invalid id parameter")
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	err = deletePublisher(id)
	if err != nil {
		err = errors.New("delete publisher err,try again ")
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/v1/publisher/list")
}

func modifyPublisherHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		idStr := c.PostForm("id")
		provinceStr := c.PostForm("province")
		cityStr := c.PostForm("city")
		nameStr := c.PostForm("name")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			msg := "invalid id parameter"
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
			return
		}
		err = modifyPublisher(id, provinceStr, cityStr, nameStr)
		if err != nil {
			msg := "modify publisher err,try again "
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/v1/publisher/list")
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
		publisherObj, err := getPublisherById(id)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid publisher id")
			return
		}
		c.HTML(http.StatusOK, "publisher/modify_publisher.html", publisherObj)
	}
}
