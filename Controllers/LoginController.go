package Controllers

import (
	"fmt"
	"net/http"

	"alexedwards.net/CC-GO/Models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	data := gin.H{
		"title": "CC - LOGIN",
	}
	c.HTML(http.StatusOK, "index.html", data)
}

func PerformLogin(c *gin.Context) {
	var user Models.User
	username := c.PostForm("username")
	password := c.PostForm("password")
	err, Id, Name, Role := Models.CheckLoginDetails(&user, username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	//Session
	/* session := sessions.Default(c)
	session.Set("id", Id)
	session.Set("name", Name)
	session.Set("role", Role) */
	fmt.Println(Id, Name)
	if Role == "super_admin" {
		AdminDashboard(c)
	} else {
		UserDashboard(c)
	}
}

func PerformLogout(c *gin.Context) {
	//Session
	session := sessions.Default(c)
	session.Set("id", "")
	session.Set("name", "")
	session.Set("role", "")
	fmt.Println("logged out successfully")
}

func UserDashboard(c *gin.Context) {
	data := gin.H{
		"title": "Normal User Login",
	}
	c.HTML(http.StatusOK, "userDashboard.html", data)
}

func AdminDashboard(c *gin.Context) {
	data := gin.H{
		"title": "Admin Login",
	}
	c.HTML(http.StatusOK, "adminDashboard.html", data)
}
