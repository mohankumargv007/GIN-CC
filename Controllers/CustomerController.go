package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerPage(c *gin.Context) {
	data := gin.H{
		"Title": "Welcome To Customer Page",
	}
	c.HTML(http.StatusOK, "CustomerLanding.tmpl", data)
}
