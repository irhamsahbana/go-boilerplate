package domain

import (
	"context"
	jwthandler "go-boilerplate/lib/jwt_handler"
)

type User struct {
	UUID				string		`bson:"uuid"`
	Name				string		`bson:"name"`
	Role				string		`bson:"role"`
	Email				string		`bson:"email"`
	Password			string		`bson:"password"`
	Phone				*string		`bson:"phone,omitempty"`
	WA					*string		`bson:"wa,omitempty"`
	ProfileUrl			*string		`bson:"profile_url,omitempty"`
	Token				*string		`bson:"token,omitempty"`
	RefreshToken		*string		`bson:"refresh_token,omitempty"`
	BlockRefreshToken	bool		`bson:"block_refresh_token"`
	CreatedAt			int64	 	`bson:"created_at"`
	UpdatedAt			*int64		`bson:"updated_at,omitempty"`
	DeletedAt			*int64		`bson:"deleted_at,omitempty"`
}

type UserUsecaseContract interface {
	RegisterUser(ctx context.Context, request *UserRegisterRequest) (*UserResponse, int, error)
	LoginUser(ctx context.Context, request *UserLoginRequest) (*UserResponse, int, error)
	FindUser(ctx context.Context, id string, withTrashed bool) (*UserResponse, int, error)
	RefreshToken(ctx context.Context, refreshToken string, claims *jwthandler.MyCustomClaims) (*UserResponse, int, error)
}

type UserRepositoryContract interface {
	FindUserBy(ctx context.Context, key string, val interface{}, withTrashed bool) (*User, int, error)
	UpdateUser(ctx context.Context, data *User) (*User, int, error)
}
