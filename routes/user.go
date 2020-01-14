package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("users")
		v1.GET("user/:id")
		v1.POST("user")
		v1.PUT("user/:id")
		v1.DELETE("user/:id")
	}
	return r
}
