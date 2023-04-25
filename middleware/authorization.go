package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/clockify/auth"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.AuthenticateToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
