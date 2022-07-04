package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example/shop-golang/service"
)

func JwtMiddleware(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer"

	authHeader := c.GetHeader("Authorization")

	if len(authHeader) <= 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	isValid, _ := service.ValidateToken(tokenString)
	if isValid {

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
