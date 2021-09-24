package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userLoginPageV2Handler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login_v2.html", nil)
}
