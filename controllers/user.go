package controllers

import (
	"MyDiaryApi/v1/env"
	auth "MyDiaryApi/v1/lib"
	"MyDiaryApi/v1/models"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	var conf models.ConfirmationPass
	c.ShouldBind(&user)
	c.ShouldBind(&conf)
	if conf.ConfPass != user.Password {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "Password tidak sesuai",
		})
		return
	}
	err := models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": err,
		})
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
	var id string
	id = c.Params.ByName("id")
	if c.MustGet("type") != nil {
		id = strconv.FormatUint(uint64(c.MustGet("id").(uint)), 10)
	}
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
		return
	}
	if err := models.LoginUser(&user, password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Password tidak cocok",
		})
		return
	} else {
		expiresIn := time.Now().Add(98 * time.Hour)
		claims := auth.Payload{
			user.ID,
			user.Email,
			"user",
			jwt.StandardClaims{
				ExpiresAt: expiresIn.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		encodedToken, err := token.SignedString([]byte(env.Environment().SecretKey))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal memproses token",
				"error":   err,
			})
			return
		}
		callback := auth.TokenResponse{
			encodedToken,
			claims,
		}
		c.JSON(http.StatusOK, callback)
	}
}
