package routes

import (
	user "MyDiaryApi/v1/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("users", user.GetAllUsers)
		v1.GET("user/:id", user.GetUser)
		v1.POST("user", user.CreateUser)
		v1.PUT("user/:id", user.UpdateUser)
		v1.DELETE("user/:id", user.DeleteUser)
		v1.POST("user/login", user.LoginUser)
	}
	return r
}
