package controllers

import (
	"MyDiaryApi/v1/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDiary(c *gin.Context) {
	var diary models.Diary
	c.ShouldBind(&diary)
	userId := c.MustGet("id").(uint)
	diary.UserID = userId
	fmt.Println("User ID : ", userId)
	fmt.Printf("This is Diary : %d", diary.UserID)
	err := models.CreateDiary(&diary)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err)
	} else {
		c.JSON(http.StatusCreated, diary)
	}
}

// func FindDiary(c *gin.Context) {
// 	var diary models.Diary
// 	c.ShouldBind(&diary)
// 	postType := c.PostForm("post_type"

// }
