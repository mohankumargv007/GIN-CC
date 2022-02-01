package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	data := gin.H{
		"title": "CC - LOGIN",
	}
	c.HTML(http.StatusOK, "index.html", data)
}
