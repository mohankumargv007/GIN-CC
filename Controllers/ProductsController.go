package Controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowProductsPage(c *gin.Context) {
	session := sessions.Default(c)
	data := gin.H{
		"Title": session.Get("name"),
	}
	c.HTML(http.StatusOK, "products.tmpl", data)
}
