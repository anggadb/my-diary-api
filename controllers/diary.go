package controllers

import (
	"MyDiaryApi/v1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDiary(c *gin.Context) {
	var diary models.Diary
	c.ShouldBind(&diary)
	userID := c.MustGet("id").(uint)
	diary.UserID = userID
	err := models.CreateDiary(&diary)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err)
	} else {
		c.JSON(http.StatusCreated, diary)
	}
}
func FindAllDiaries(c *gin.Context) {
	var diaries []models.Diary
	userID := c.MustGet("id").(uint)
	models.FindAllDiaries(&diaries, userID)
	c.JSON(http.StatusOK, diaries)
}
func UpdateDiary(c *gin.Context) {
	var diary models.Diary
	id := c.Params.ByName("id")
	c.ShouldBind(&diary)
	err := models.UpdateDiary(&diary, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, err)
	} else {
		c.JSON(http.StatusOK, diary)
	}
}
