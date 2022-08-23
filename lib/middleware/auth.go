package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	// clientToken := c.Request.Header.Get("token")

	// if clientToken == "" {
	// 	http_response.ReturnResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
	// 	c.Abort()
	// 	return
	// }

	// tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
}
