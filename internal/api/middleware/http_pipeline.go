package middleware

import (
	"github.com/gin-gonic/gin"
)

func AddHttpMiddleware(api *gin.Engine) {
	api.Use(tokenAuthorization())
}

func errorResponse(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
