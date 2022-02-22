package Routes

import (
	"alexedwards.net/CC-GO/Controllers"
	"alexedwards.net/CC-GO/Middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.Use(Middlewares.AuthMiddleware)
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}

	r.GET("login", Controllers.ShowLoginPage)
	r.POST("login", Controllers.PerformLogin)

	r.GET("logout", Controllers.PerformLogout)

	v2 := r.Group("/v2")
	v2.Use(Middlewares.AuthMiddleware)
	{
		users := v2.Group("users")
		{
			users.POST("create_super_user/", Controllers.CreateAUser)
			//users.POST("login/", Controllers.login)
		}
	}

	v3 := r.Group("/user")
	{
		v3.GET("userDashboard/", Controllers.UserDashboard)
	}

	v4 := r.Group("/admin")
	{
		v4.GET("adminDashboard/", Controllers.AdminDashboard)
	}
	return r
}
