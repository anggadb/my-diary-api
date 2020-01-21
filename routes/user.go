package routes

import (
	user "MyDiaryApi/v1/controllers"
	"github.com/gin-gonic/gin"
	"os"
)

func UserRouter(route *gin.Engine) {
	v := route.Group(os.Getenv("API_VERSION"))
	v.POST("user", user.CreateUser)
	v.GET("users", user.GetAllUsers)
	v.GET("user/:id", user.GetUser)
	v.PUT("user/:id", user.UpdateUser)
	v.DELETE("user/:id", user.DeleteUser)
	v.POST("user/login", user.LoginUser)
}
