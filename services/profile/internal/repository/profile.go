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
	GetMany(ctx context.Context, ids []string) ([]datastruct.Profile, error)
	GetById(ctx context.Context, id string) (datastruct.Profile, error)
	GetByUsername(ctx context.Context, username string) (datastruct.Profile, error)
	UsernameTaken(ctx context.Context, username string) bool
}

type profileQuery struct {
	col *mongo.Collection
}

func (pq *profileQuery) Create(ctx context.Context, u datastruct.Profile) (datastruct.Profile, error) {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return datastruct.Profile{}, err
	}

	res, err := pq.
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

func (pq *profileQuery) GetById(ctx context.Context, id string) (res datastruct.Profile, err error) {
	err = pq.
		col.
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (pq *profileQuery) GetMany(ctx context.Context, ids []string) (res []datastruct.Profile, err error) {
	filter := bson.M{"_id": bson.M{"$in": ids}}

	cur, err := pq.col.Find(ctx, filter)
	if err != nil {
		return res, err
	}

	err = cur.All(ctx, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (pq *profileQuery) Delete(ctx context.Context, id string) error {
	res, err := pq.
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

func (pq *profileQuery) Update(ctx context.Context, u datastruct.Profile) (res datastruct.Profile, err error) {
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

	err = pq.
		col.
		FindOneAndUpdate(ctx, filter, input, &opt).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (pq *profileQuery) GetByUsername(ctx context.Context, username string) (res datastruct.Profile, err error) {
	err = pq.
		col.
		FindOne(ctx, bson.M{"username": username}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (pq *profileQuery) UsernameTaken(ctx context.Context, username string) bool {
	user := datastruct.Profile{}
	err := pq.col.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return false
	}

	if user.Id.Hex() == "" {
		return false
	}

	return true
}
