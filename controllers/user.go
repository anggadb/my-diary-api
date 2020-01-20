package controllers

import (
	auth "MyDiaryApi/v1/lib"
	"MyDiaryApi/v1/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	err := models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err)
	} else {
		c.JSON(http.StatusCreated, user)
	}

}
func GetAllUsers(c *gin.Context) {
	var u []models.User
	models.FindAllUsers(&u)
	c.JSON(http.StatusOK, u)
}
func GetUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.FindUserById(&user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.FindUserById(&user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	}
	c.ShouldBind(&user)
	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Berhasil menghapus user dengan ID : " + id,
		})
	}
}
func LoginUser(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	email := user.Email
	password := user.Password
	err := models.FindUserByEmail(&user, email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, err)
	}
	if err := models.LoginUser(&user, password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Password tidak cocok",
		})
	} else {
		claims := auth.Payload{
			user.ID,
			user.Email,
			jwt.StandardClaims{ExpiresAt: 15000},
		}
		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		token, err := sign.SignedString([]byte(os.Getenv("SECRET_KEY")))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusProcessing, gin.H{
				"messager": "Gagal memproses login",
			})
		}
		callback := auth.TokenResponse{
			token,
			auth.Payload{
				Email: claims.Email,
				ID:    claims.ID,
			},
		}
		c.JSON(http.StatusOK, callback)
	}
}
