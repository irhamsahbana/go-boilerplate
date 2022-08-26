package mongo

import (
	"context"
	"go-boilerplate/domain"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *userRepository) UpdateUser(ctx context.Context, data *domain.User) (*domain.User, int, error) {
	var user domain.User

	filter := bson.M{"uuid": data.UUID}
	opts := options.Update().SetUpsert(true)

	updateFields := bson.M{
		"$set": bson.M{
			"token": data.Token,
			"refresh_token": data.RefreshToken,
			"block_refresh_token": data.BlockRefreshToken,
			"updated_at": time.Now().UTC().UnixMicro(),
		},
	}

	if _, err := repo.Collection.UpdateOne(ctx, filter, updateFields, opts); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}
