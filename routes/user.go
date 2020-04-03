package routes

import (
	user "MyDiaryApi/v1/controllers"
	"MyDiaryApi/v1/env"
	lib "MyDiaryApi/v1/lib"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.Engine) {
	v := route.Group(env.Environment().APIVersion)
	v.POST("user", user.CreateUser)
	v.POST("user/login", user.LoginUser)
	v.GET("user", lib.Auth, user.GetUser)
	// ADMIN
	v.GET("users", user.GetAllUsers)
	v.GET("user/:id", user.GetUser)
	v.PUT("user/:id", user.UpdateUser)
	v.DELETE("user/:id", user.DeleteUser)
}
