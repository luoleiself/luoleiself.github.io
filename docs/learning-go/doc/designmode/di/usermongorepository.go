package di

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	db *mongo.Collection
}

func NewUserMongoRepository(db *mongo.Collection) *UserMongoRepository {
	return &UserMongoRepository{db: db}
}

func (r *UserMongoRepository) Save(ctx context.Context, user *User) error {
	_, err := r.db.InsertOne(ctx, bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		return err
	}
	return nil
}
