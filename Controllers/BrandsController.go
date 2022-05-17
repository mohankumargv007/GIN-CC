package Controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowBrandsPage(c *gin.Context) {
	session := sessions.Default(c)
	data := gin.H{
		"Title": session.Get("name"),
	}
	c.HTML(http.StatusOK, "brands.tmpl", data)
}
