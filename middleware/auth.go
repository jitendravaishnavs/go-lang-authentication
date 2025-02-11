package middleware

import (
	"authentication/helpers"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFun {
	return func(c *gin.Context) {
		authHeaders := c.GetHeader("Authorization")
		if authHeaders == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found"})
			c.Abort()
			return
		}

		//Bearer <token>
		authHeader = strings.TrimPrefix(authHeader, "Bearer")

		claims, err := helpers.ValidateToken(authHeader)

		if err != nil {
			log.Printf("Token Validation Error:%v", err)

			c.JSON(http.StatusUnauthorized, gin.H{"error": "InValid Token"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}

}
