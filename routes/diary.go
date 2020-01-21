package routes

import (
	diary "MyDiaryApi/v1/controllers"
	"MyDiaryApi/v1/lib"
	"github.com/gin-gonic/gin"
	"os"
)

func DiaryRouter(route *gin.Engine) {
	v := route.Group(os.Getenv("API_VERSION"))
	v.POST("diary", lib.Auth, diary.CreateDiary)
	v.GET("diaries", lib.Auth, diary.FindAllDiaries)
	v.PUT("diary/:id", lib.Auth, diary.UpdateDiary)
}
