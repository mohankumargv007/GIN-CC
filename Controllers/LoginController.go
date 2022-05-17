package Controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"alexedwards.net/CC-GO/Models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Username string
	Password string
}

type Credentials struct {
	user     Models.User
	username string
	password string
}

func ShowLoginPage(c *gin.Context) {
	data := gin.H{
		"title": "CC - LOGIN",
	}
	session := sessions.Default(c)
	user := session.Get("id")
	if user != nil {
		navigateToRoleBasedAccess(c)
		return
	}
	c.HTML(http.StatusOK, "index.html", data)
}

func PerformLogin(c *gin.Context) {
	var user Models.User
	var requestBody RequestBody
	var Credentials Models.LoginCredentials

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "Please give proper credentails"})
	}

	fmt.Println(user)

	Credentials.Username = requestBody.Username
	Credentials.Password = requestBody.Password

	fmt.Println(requestBody)

	err, UserData := Models.CheckLoginDetails(Credentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "Username & Passoword doesn't match.Please give proper values"})
		return
	}

	//Session
	session := sessions.Default(c)
	session.Set("id", UserData.ID)
	session.Set("name", UserData.FirstName)
	session.Set("role", UserData.Role)
	session.Save()

	navigateToRoleBasedAccess(c)
}

func PerformLogout(c *gin.Context) {
	//Session
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	fmt.Println("logged out successfully")
	location := url.URL{Path: "/login"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func navigateToRoleBasedAccess(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("role")
	if role == "super_admin" {
		location := url.URL{Path: "/admin/adminDashboard"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
	if role == "user" {
		location := url.URL{Path: "/user/userDashboard"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func UserDashboard(c *gin.Context) {
	data := gin.H{
		"title": "Normal User Login",
	}
	c.HTML(http.StatusOK, "userDashboard.tmpl", data)
}

func AdminDashboard(c *gin.Context) {
	session := sessions.Default(c)
	data := gin.H{
		"Title": session.Get("name"),
	}
	c.HTML(http.StatusOK, "adminDashboard.tmpl", data)
}
