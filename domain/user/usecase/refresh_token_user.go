package usecase

import (
	"context"
	"errors"
	"go-boilerplate/domain"
	jwthandler "go-boilerplate/lib/jwt_handler"
	"net/http"
)

func (u *userUsecase) RefreshToken(c context.Context, refreshToken string, claims *jwthandler.MyCustomClaims) (*domain.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	var currentRefreshToken string

	result, code, err := u.userRepo.FindUserBy(ctx, "uuid", claims.UUID, false)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	if result.RefreshToken != nil {
		currentRefreshToken = *result.RefreshToken
	}

	if currentRefreshToken != refreshToken || result.BlockRefreshToken == true {
		return nil, http.StatusUnauthorized, errors.New("Refresh token is invalid")
	}

	token, refreshToken, err := jwthandler.GenerateAllTokens(result.UUID, result.Name, result.Role)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	result.Token = &token
	result.RefreshToken = &refreshToken

	result, code, err  = u.userRepo.UpdateUser(ctx, result)
	if err != nil {
		return nil, code, err
	}

	var resp domain.UserResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.Role = result.Role
	resp.Token = &token
	resp.RefreshToken = &refreshToken

	return &resp, http.StatusOK, nil
}
