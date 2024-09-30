package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ssoql/auth-service/internal/api"
)

func tokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI != "/token" {
			token := c.Request.Header.Get("Authorization")

			if token == "" {
				errorResponse(c, http.StatusUnauthorized, "API token required")
				c.Abort()
				return
			}

			// Expecting "Bearer <token>"
			tokenParts := strings.Split(token, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				errorResponse(c, http.StatusUnauthorized, "Invalid Authorization header format")
				c.Abort()
				return
			}

			if !api.IsTokenValid(tokenParts[1]) {
				errorResponse(c, http.StatusUnauthorized, "invalid API token")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
