package Controllers

import (
	"fmt"
	"net/http"

	"alexedwards.net/CC-GO/Models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateAUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	user.Password = "Mohan@123"
	password := user.Password
	user.Password, _ = HashPassword(password)
	fmt.Println(user.Password)
	fmt.Println(string(user.Role))
	err := Models.CreateAUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
