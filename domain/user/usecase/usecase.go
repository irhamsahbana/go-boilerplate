package usecase

import (
	"go-boilerplate/domain"
	"time"
)

type userUsecase struct {
	userRepo domain.UserRepositoryContract
	contextTimeout time.Duration
}

func NewUserUsecase(repo domain.UserRepositoryContract, timeout time.Duration) domain.UserUsecaseContract {
	return &userUsecase{
		userRepo: repo,
		contextTimeout: timeout,
	}
}
