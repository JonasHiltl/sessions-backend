package repository

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dao interface {
	NewProfileRepository() ProfileQuery
}

type dao struct {
	db *mongo.Database
}

const (
	PROFILE_COLLECTION = "profiles"
)

func NewDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url, exists := os.LookupEnv("MONGO_URL")
	if !exists {
		return nil, errors.New("Mongo url not defined")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	db := client.Database("sessions")
	models := []mongo.IndexModel{
		{
			Keys: bson.M{"username": 1}, Options: options.Index().SetUnique(true),
		},
	}

	db.Collection(PROFILE_COLLECTION).Indexes().CreateMany(ctx, models)

	return db, nil
}

func NewDAO(db *mongo.Database) Dao {
	return &dao{db: db}
}

func (d *dao) NewProfileRepository() ProfileQuery {
	return &profileQuery{col: d.db.Collection(PROFILE_COLLECTION)}
}
