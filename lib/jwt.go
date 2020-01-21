package lib

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Payload struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Type  string `json:"type"`
	jwt.StandardClaims
}
type TokenResponse struct {
	Token string `json:"token"`
	Payload
}

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("authorization")
	claims := &Payload{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Gagal memverifikasi algoritma token",
			})
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Gagal memproses token",
			"error":   err,
		})
	}
	if !parsedToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token tidak valid",
		})
	}
	c.Set("id", claims.ID)
	c.Set("type", claims.Type)
	c.Next()
}
