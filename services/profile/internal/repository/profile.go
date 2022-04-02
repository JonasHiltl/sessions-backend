package repository

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProfileQuery interface {
	Create(ctx context.Context, u datastruct.Profile) (datastruct.Profile, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, u datastruct.Profile) (datastruct.Profile, error)
	GetById(ctx context.Context, id string) (datastruct.Profile, error)
	GetByUsername(ctx context.Context, username string) (datastruct.Profile, error)
	UsernameTaken(ctx context.Context, username string) bool
}

type profileQuery struct {
	col *mongo.Collection
}

func (uq *profileQuery) Create(ctx context.Context, u datastruct.Profile) (datastruct.Profile, error) {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return datastruct.Profile{}, err
	}

	res, err := uq.
		col.
		InsertOne(ctx, u)
	if err != nil {
		return datastruct.Profile{}, err
	}

	id := res.InsertedID.(primitive.ObjectID)

	if res.InsertedID != nil {
		u.Id = id
	}

	return u, nil
}

func (uq *profileQuery) GetById(ctx context.Context, id string) (res datastruct.Profile, err error) {
	err = uq.
		col.
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uq *profileQuery) Delete(ctx context.Context, id string) error {
	res, err := uq.
		col.
		DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("Failed to delete user")
	}

	return nil
}

func (uq *profileQuery) Update(ctx context.Context, u datastruct.Profile) (res datastruct.Profile, err error) {
	input := bson.M{}
	filter := bson.M{"_id": u.Id}

	after := options.After
	opt := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	if u.Username != "" {
		input["username"] = u.Username
	}

	if u.Firstname != "" {
		input["firstname"] = u.Firstname
	}

	if u.Lastname != "" {
		input["lastname"] = u.Lastname
	}

	if u.Avatar != "" {
		input["avatar"] = u.Avatar
	}

	err = uq.
		col.
		FindOneAndUpdate(ctx, filter, input, &opt).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uq *profileQuery) GetByUsername(ctx context.Context, username string) (res datastruct.Profile, err error) {
	err = uq.
		col.
		FindOne(ctx, bson.M{"username": username}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uq *profileQuery) UsernameTaken(ctx context.Context, username string) bool {
	user := datastruct.Profile{}
	err := uq.col.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return false
	}

	if user.Id.Hex() == "" {
		return false
	}

	return true
}
