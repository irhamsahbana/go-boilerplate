package http

import (
	"context"
	"go-boilerplate/domain"
	"go-boilerplate/lib/http_response"
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

	router.GET("users", handler.Profile)

	router.Use(middleware.Auth)
	router.POST("users/register", handler.Register)
	router.POST("users/login", handler.Login)
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
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
	}

	http_response.ReturnResponse(c, httpCode, "Registred", result)
}

func (h *UserHandler) Login(c *gin.Context) {
	http_response.ReturnResponse(c, http.StatusOK, "Logged", nil)
}

func (h *UserHandler) Profile(c *gin.Context) {
	http_response.ReturnResponse(c, http.StatusOK, "you are ...", nil)
}
