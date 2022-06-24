package middleware

import (
	"graphics-app-backend/src/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")
		if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
		} else {
			c.String(http.StatusBadRequest, "Unauthorized")
			c.Abort()
			return
		}
		claims, err := helpers.AuthenticateGoogle(token)
		if err != nil {
			c.String(http.StatusUnauthorized, "Error while authenticating with Google")
			c.Abort()
			return
		}
		userId := claims["sub"].(string)
		c.Set("userId", userId)
		c.Next()
	}
}
