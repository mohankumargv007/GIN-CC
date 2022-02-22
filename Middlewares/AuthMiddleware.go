package Middlewares

import (
	"github.com/gin-gonic/gin"
)

/* func AuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"mohankumar": "secret",
	})
} */

func AuthMiddleware(c *gin.Context) {
	/* session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} */
}

/* func AuthMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	}
} */
