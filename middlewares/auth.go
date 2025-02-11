package middleware

import (
	"context"
	"net/http"
	"strings"
	"user-authentication/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID, err := utils.ValidateToken(bearerToken[1], jwtSecret)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(c.Request.Context(), "userID", userID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
