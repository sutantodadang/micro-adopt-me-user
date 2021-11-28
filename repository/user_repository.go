package repository

import (
	"context"
	"user_service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(user models.RegisterUser) error
	Find(email string) (models.RegisterUser, error)
}

type UserRepo struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(user models.RegisterUser) error {
	_, err := r.db.Collection("user").InsertOne(context.Background(), &user)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Find(email string) (models.RegisterUser, error) {
	res := r.db.Collection("user").FindOne(context.Background(), bson.M{"email": email})

	var result models.RegisterUser

	if err := res.Decode(&result); err != nil {
		return result, err
	}

	return result, nil

}
