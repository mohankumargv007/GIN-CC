package Routes

import (
	"alexedwards.net/CC-GO/Controllers"
	"alexedwards.net/CC-GO/Middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", Controllers.ShowLoginPage)

	v1 := r.Group("/v1")
	v1.Use(Middlewares.AuthMiddleware())
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}

	v2 := r.Group("/v2")
	{
		users := v2.Group("users")
		{
			users.POST("create_super_user/", Controllers.CreateAUser)
		}
	}

	return r
}
