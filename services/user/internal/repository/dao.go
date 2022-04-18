package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Dao interface {
	NewUserRepository() UserRepository
	NewProfileRepository() ProfileRepository
}

type dao struct {
	db *mongo.Database
}

const (
	USER_COLLECTION = "user"
)

func NewDB(url string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	db := client.Database("sessions")

	models := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "email", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bsonx.Doc{{Key: "username", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true)},
	}

	_, err = db.Collection(USER_COLLECTION).Indexes().CreateMany(ctx, models)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDAO(db *mongo.Database) Dao {
	return &dao{db: db}
}

func (d *dao) NewUserRepository() UserRepository {
	return &userRepository{col: d.db.Collection(USER_COLLECTION)}
}

func (d *dao) NewProfileRepository() ProfileRepository {
	return &profileRepository{col: d.db.Collection(USER_COLLECTION)}
}
