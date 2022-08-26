package middleware

import (
	"fmt"
	"go-boilerplate/lib/http_response"
	jwthandler "go-boilerplate/lib/jwt_handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		http_response.ReturnResponse(c, http.StatusUnauthorized, http.StatusText(401), nil)
		c.Abort()
		return
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]

	fmt.Println(tokenString)

	claims, err := jwthandler.ValidateToken(tokenString)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}

	fmt.Println(claims)

	c.Next()
}
