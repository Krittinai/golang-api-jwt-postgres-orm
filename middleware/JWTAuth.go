package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthen() gin.HandlerFunc {
	return func(c *gin.Context) {
		mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Unauthorized"})
			c.Abort()
		}
		tokenString := strings.Split(header, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return mySigningKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Set example variable
			fmt.Print(claims["userId"])
			c.Set("uesrId", "xxxx-xxxx-xxxx-xxxx-xxxx-xxxx-xxxx-xxxx")
		} else {
			c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": err.Error()})
			c.Abort()
		}

		c.Next()
	}
}
