package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dao interface {
	NewAuthRepository() AuthQuery
}

type dao struct {
	db *mongo.Database
}

const (
	AUTH_COLLECTION = "auth"
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
			Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true),
		},
	}

	db.Collection(AUTH_COLLECTION).Indexes().CreateMany(ctx, models)

	return db, nil
}

func NewDAO(db *mongo.Database) Dao {
	return &dao{db: db}
}

func (d *dao) NewAuthRepository() AuthQuery {
	return &authQuery{col: d.db.Collection(AUTH_COLLECTION)}
}
