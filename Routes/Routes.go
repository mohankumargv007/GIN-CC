package Routes

import (
	"alexedwards.net/CC-GO/Controllers"
	"alexedwards.net/CC-GO/Middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.Use(static.Serve("/CSS", static.LocalFile("./UI/Statics/CSS", true)))
	r.Use(static.Serve("/JS", static.LocalFile("./UI/Statics/JS", true)))

	v1 := r.Group("/v1")
	v1.Use(Middlewares.AuthMiddleware)
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}

	r.GET("/", Controllers.CustomerPage)

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
		v4.GET("brands/", Controllers.ShowBrandsPage)
		v4.GET("products/", Controllers.ShowProductsPage)
	}

	v5 := r.Group("/v5")
	//v5.Use(Middlewares.AuthMiddleware)
	{
		v5.GET("category", Controllers.GetAllCategories)
		v5.POST("createCatergory", Controllers.CreateCategory)
		v5.GET("category/:id", Controllers.GetACategory)
		v5.PUT("category/:id", Controllers.UpdateACategory)
		v5.DELETE("category/:id", Controllers.DeleteACategory)
	}

	return r
}
