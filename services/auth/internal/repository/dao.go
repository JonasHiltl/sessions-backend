package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dao interface {
	NewAuthRepository() AuthRepository
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
	mod := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		}, Options: options.Index().SetUnique(true),
	}

	_, err = db.Collection(AUTH_COLLECTION).Indexes().CreateOne(ctx, mod)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDAO(db *mongo.Database) Dao {
	return &dao{db: db}
}

func (d *dao) NewAuthRepository() AuthRepository {
	return &authRepository{col: d.db.Collection(AUTH_COLLECTION)}
}
