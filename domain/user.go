package domain

import "context"

type User struct {
	UUID			string		`bson:"uuid"`
	FirstName		string		`bson:"first_name"`
	LastName		string		`bson:"last_name"`
	Password		string		`bson:"password"`
	Email			string		`bson:"email"`
	Phone			*string		`bson:"phone,omitempty"`
	Type			string		`bson:"type"`
	Token			*string		`bson:"token,omitempty"`
	RefreshToken	*string		`bson:"refresh_token,omitempty"`
	CreatedAt		int64	 	`bson:"created_at"`
	UpdatedAt		*int64		`bson:"updated_at,omitempty"`
	DeletedAt		*int64		`bson:"deleted_at,omitempty"`
}

type UserUsecaseContract interface {
	RegisterUser(ctx context.Context, request *UserRegisterRequest) (*UserResponse, int, error)
	LoginUser(ctx context.Context, request *UserLoginRequest) (*UserResponse, int, error)
}

type UserRepositoryContract interface {

}
