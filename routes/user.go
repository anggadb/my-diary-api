package routes

import (
	user "MyDiaryApi/v1/controllers"
	lib "MyDiaryApi/v1/lib"
	"os"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.Engine) {
	v := route.Group(os.Getenv("API_VERSION"))
	v.POST("user", user.CreateUser)
	v.POST("user/login", user.LoginUser)
	v.GET("user", lib.Auth, user.GetUser)
	// ADMIN
	v.GET("users", user.GetAllUsers)
	v.GET("user/:id", user.GetUser)
	v.PUT("user/:id", user.UpdateUser)
	v.DELETE("user/:id", user.DeleteUser)
}
