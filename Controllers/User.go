package Controllers

import (
	"net/http"

	"alexedwards.net/CC-GO/Models"
	"github.com/gin-gonic/gin"
)

func CreateAUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateAUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
