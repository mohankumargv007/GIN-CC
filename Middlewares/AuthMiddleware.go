package Middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/* func AuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"mohankumar": "secret",
	})
} */

func AuthMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("id")
	if user == nil {
		data := gin.H{
			"title": "CC - LOGIN",
		}
		c.HTML(http.StatusOK, "index.html", data)
	}
}
