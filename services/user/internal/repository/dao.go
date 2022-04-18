package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	mod := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		}, Options: options.Index().SetUnique(true),
	}

	_, err = db.Collection(USER_COLLECTION).Indexes().CreateOne(ctx, mod)
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
