package usecase

import (
	"context"
	"errors"
	"go-boilerplate/domain"
	jwthandler "go-boilerplate/lib/jwt_handler"
	passwordhandler "go-boilerplate/lib/password_handler"
	"net/http"
)

func (u *userUsecase) LoginUser(c context.Context, req *domain.UserLoginRequest) (*domain.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.userRepo.FindUserBy(ctx, "email", req.Email, false)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	if ok := passwordhandler.VerifyPassword(result.Password, req.Password); !ok {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	token, refreshtoken, err := jwthandler.GenerateAllTokens(result.UUID, result.Name, result.Role)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var resp domain.UserResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.Role = result.Role
	resp.Token = &token
	resp.RefreshToken = &refreshtoken

	return &resp, http.StatusOK, nil
}
