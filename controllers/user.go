package controllers

import (
	"MyDiaryApi/v1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSON struct {
	message string
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusConflict)
	} else {
		c.JSON(http.StatusCreated, user)
	}

}
func GetAllUsers(c *gin.Context) {
	var u []models.User
	err := models.FindAllUsers(&u)
	if err != nil {
		c.AbortWithStatus(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, u)
	}
}
func GetUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.FindUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.FindUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusConflict)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	message := JSON{message: "Berhasil menghapus user"}
	error := JSON{message: "Gagal menghapus user"}
	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, error)
	} else {
		c.JSON(http.StatusOK, message)
	}
}
