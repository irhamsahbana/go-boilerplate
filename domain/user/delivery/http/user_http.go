package http

import (
	"context"
	"go-boilerplate/domain"
	"go-boilerplate/lib/http_response"
	jwthandler "go-boilerplate/lib/jwt_handler"
	"go-boilerplate/lib/middleware"
	"net/http"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase domain.UserUsecaseContract
}

func NewUserHandler(router *gin.Engine, usecase domain.UserUsecaseContract) {
	handler := &UserHandler{
		UserUsecase: usecase,
	}

	authorized := router.Group("/", middleware.Auth)
	authorized.GET("/auth/logout", handler.Logout)
	authorized.GET("/users/profile", handler.Profile)

	router.POST("auth/login", handler.Login)
	// router.POST("auth/register", handler.Register)
	router.GET("auth/refresh-token", handler.RefreshToken)
}

// func (h *UserHandler) Register(c *gin.Context) {
// 	var request domain.UserRegisterRequest

// 	err := c.BindJSON(&request)
// 	if err != nil {
// 		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
// 		return
// 	}

// 	ctx := context.Background()
// 	result, httpCode, err := h.UserUsecase.RegisterUser(ctx, &request)
// 	if err != nil {
// 		http_response.ReturnResponse(c, httpCode, err.Error(), err.Error())
// 	}

// 	http_response.ReturnResponse(c, httpCode, "Registred", result)
// }

func (h *UserHandler) Login(c *gin.Context) {
	var request domain.UserLoginRequest

	if err := c.BindJSON(&request); err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.Login(ctx, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		c.Abort()
		return
	}

	http_response.ReturnResponse(c, httpCode, "Authenticated", result)
}

func (h *UserHandler) Profile(c *gin.Context) {
	http_response.ReturnResponse(c, http.StatusOK, "you are ...", nil)
}

func (h *UserHandler) RefreshToken(c *gin.Context)  {
	accessToken := c.GetHeader("X-ACCESS-TOKEN")
	refreshToken := c.GetHeader("X-REFRESH-TOKEN")

	_, err := jwthandler.ValidateToken(accessToken)
	if err != nil {
		v, _ := err.(*jwt.ValidationError)

		if v.Errors == jwt.ValidationErrorExpired {
		} else {
			http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
			return
		}
	}

	claimsRT, err := jwthandler.ValidateToken(refreshToken)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.RefreshToken(ctx, accessToken, refreshToken, claimsRT.UserUUID)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		c.Abort()
		return
	}

	http_response.ReturnResponse(c, httpCode, "Token Refreshed", result)
}

func (h *UserHandler) Logout(c *gin.Context) {
	AT := c.GetString("access_token")
	userId := c.GetString("user_uuid")

	ctx := context.Background()
	_, httpcode, err := h.UserUsecase.Logout(ctx, AT, userId)
	if err != nil {
		http_response.ReturnResponse(c, httpcode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpcode, "Logout", nil)
}
