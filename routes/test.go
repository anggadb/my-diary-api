package routes

import (
	"MyDiaryApi/v1/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestRouter(route *gin.Engine) {
	v := route.Group(env.Environment().APIVersion)
	v.GET("test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bisa",
		})
	})
}
