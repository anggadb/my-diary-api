package routes

import (
	diary "MyDiaryApi/v1/controllers"
	"MyDiaryApi/v1/env"
	"MyDiaryApi/v1/lib"

	"github.com/gin-gonic/gin"
)

func DiaryRouter(route *gin.Engine) {
	v := route.Group(env.Environment().APIVersion)
	v.POST("diary", lib.Auth, diary.CreateDiary)
	v.GET("diaries", lib.Auth, diary.FindAllDiaries)
	v.PUT("diary/:id", lib.Auth, diary.UpdateDiary)
}
