package router

import (
	"chris_project/controllers"
	"chris_project/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//初始化一個基於redis的session store
	user := r.Group("/user", middlewares.EnableRedisSession())
	{
		//註冊用戶
		user.POST("/register", controllers.Register)
		//登入用戶
		user.POST("/login", controllers.Login)
		//登出用戶
		user.GET("/logout", controllers.Logout)

		//檢查http request裡的cookie有沒有sessionID
		authorized := user.Group("/auth", middlewares.AuthSessionMiddle())
		{
			//知道目前登入的是哪位使用者
			authorized.GET("/current_user", controllers.CurrentUser)

			//透過param取得指定product的ID
			authorized.GET("/product/:pid", controllers.GetProduct)
			//透過JSON創建product
			authorized.POST("/product", controllers.CreateProduct)
			//透過param更新指定product的ID
			//product的資料用JSON檔傳送
			authorized.PUT("/product/:pid", controllers.UpdateProduct)
			//透過param刪除指定product的ID
			authorized.DELETE("/product/:pid", controllers.DeleteProduct)
		}

	}

	return r
}
