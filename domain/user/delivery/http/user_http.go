package http

import (
	"context"
	"go-boilerplate/domain"
	"go-boilerplate/lib/http_response"
	jwthandler "go-boilerplate/lib/jwt_handler"
	"go-boilerplate/lib/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase domain.UserUsecaseContract
}

func NewUserHandler(router *gin.Engine, usecase domain.UserUsecaseContract) {
	handler := &UserHandler{
		UserUsecase: usecase,
	}

	router.POST("users/login", handler.Login)
	router.POST("users/register", handler.Register)
	router.GET("users/refresh-token", handler.RefreshToken)

	router.Use(middleware.Auth)
	router.GET("users/profile", handler.Profile)
}

func (h *UserHandler) Register(c *gin.Context) {
	var request domain.UserRegisterRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.RegisterUser(ctx, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), err.Error())
	}

	http_response.ReturnResponse(c, httpCode, "Registred", result)
}

func (h *UserHandler) Login(c *gin.Context) {
	var request domain.UserLoginRequest

	if err := c.BindJSON(&request); err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.LoginUser(ctx, &request)
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
	cookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}

	refreshToken := cookie.Value

	claims, err := jwthandler.ValidateToken(refreshToken)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.RefreshToken(ctx, refreshToken, claims)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		c.Abort()
		return
	}

	c.SetCookie("refresh_token", *result.RefreshToken, 3600, "", "", false, true)
	result.RefreshToken = nil
	http_response.ReturnResponse(c, httpCode, "Token Refreshed", result)
}
