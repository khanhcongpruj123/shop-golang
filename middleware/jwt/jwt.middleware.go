package jwt

import (
	"fmt"
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
	isValid, err := service.ValidateToken(tokenString)
	if isValid {

	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
