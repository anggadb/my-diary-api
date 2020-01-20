package lib

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Payload struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
type TokenResponse struct {
	Token string `json:"token"`
	Payload
}

func auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != t.Method {
			return nil, fmt.Errorf("Signing method tidak diterima : %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if parsedToken == nil && err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Tidak mempunyai akses",
		})
	} else {
		payload := parsedToken.Claims.(jwt.MapClaims)
		fmt.Printf("ID : %s", payload["ID"])
		c.Set("id", payload["ID"])
		c.Next()
	}
}
