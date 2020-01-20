package controllers

import (
	"MyDiaryApi/v1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDiary(c *gin.Context) {
	var diary models.Diary
	c.ShouldBind(&diary)
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
// 	postType := c.PostForm("post_type")
// 	fmt.Printf("Post type : %s", postType)
// 	if postType == "user" {
// 		err := models.FindDiaryByUser()
// 	}
// }
