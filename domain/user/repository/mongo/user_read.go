package mongo

import (
	"go-boilerplate/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	DB			mongo.Database
	Collection	mongo.Collection
}

func NewUserMongoRepository(DB mongo.Database) domain.UserRepositoryContract {
	return &userRepository{
		DB: DB,
		Collection: *DB.Collection("users"),
	}
}
