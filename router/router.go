package router

import (
	"chris_project/controllers"
	"chris_project/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user", middlewares.EnableRedisSession())
	{
		user.POST("/register", controllers.Register)
		user.POST("/login", controllers.Login)
		user.GET("/logout", controllers.Logout)

		authorized := user.Group("/auth", middlewares.AuthSessionMiddle())
		{
			authorized.GET("/current_user", controllers.CurrentUser)

			authorized.GET("/product/:pid", controllers.GetProduct)
			authorized.POST("/product", controllers.CreateProduct)
			authorized.PUT("/product/:pid", controllers.UpdateProduct)
			authorized.DELETE("/product/:pid", controllers.DeleteProduct)
		}

	}

	return r
}
