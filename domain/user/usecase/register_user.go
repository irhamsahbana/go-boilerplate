package usecase

import (
	"context"
	"go-boilerplate/domain"
)

func (u *userUsecase) RegisterUser(c context.Context, req *domain.UserRegisterRequest) (*domain.UserResponse, int, error) {
	return nil, 409, nil
}

func (u *userUsecase) LoginUser(c context.Context, req *domain.UserLoginRequest) (*domain.UserResponse, int, error) {
	return nil, 200, nil
}

func (u *userUsecase) ProfileUser(c context.Context, authUser *domain.User) (*domain.UserResponse, int, error) {
	return nil, 200, nil
}

// helper function

func hashPassword() {
	// some = "string"

	// bcrypt.
}
