package repository

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserQuery interface {
	Create(ctx context.Context, u datastruct.User) (datastruct.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, u datastruct.User) (datastruct.User, error)
	GetById(ctx context.Context, id string) (datastruct.User, error)
	GetByEmail(ctx context.Context, email string) (datastruct.User, error)
	GetByUsername(ctx context.Context, username string) (datastruct.User, error)
	UsernameTaken(ctx context.Context, username string) bool
}

type userQuery struct {
	col *mongo.Collection
}

func (uq *userQuery) Create(ctx context.Context, u datastruct.User) (datastruct.User, error) {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return datastruct.User{}, err
	}

	res, err := uq.
		col.
		InsertOne(ctx, u)
	if err != nil {
		return datastruct.User{}, err
	}

	id := res.InsertedID.(primitive.ObjectID)

	if res.InsertedID != nil {
		u.Id = id
	}

	return u, nil
}

func (uq *userQuery) GetById(ctx context.Context, id string) (res datastruct.User, err error) {
	err = uq.
		col.
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uq *userQuery) Delete(ctx context.Context, id string) error {
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

func (uq *userQuery) Update(ctx context.Context, u datastruct.User) (res datastruct.User, err error) {
	input := bson.M{}
	filter := bson.M{"_id": u.Id}

	after := options.After
	opt := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	if u.Username != "" {
		input["username"] = u.Username
	}

	if u.Email != "" {
		input["email"] = u.Email
	}

	if u.Firstname != "" {
		input["firstname"] = u.Firstname
	}

	if u.Lastname != "" {
		input["lastname"] = u.Lastname
	}

	if u.Password != "" {
		input["password"] = u.Password
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

func (uq *userQuery) GetByEmail(ctx context.Context, email string) (res datastruct.User, err error) {
	err = uq.
		col.
		FindOne(ctx, bson.M{"email": email}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uq *userQuery) GetByUsername(ctx context.Context, username string) (res datastruct.User, err error) {
	err = uq.
		col.
		FindOne(ctx, bson.M{"username": username}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uq *userQuery) UsernameTaken(ctx context.Context, username string) bool {
	user := datastruct.User{}
	err := uq.col.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return false
	}

	if user.Id.Hex() == "" {
		return false
	}

	return true
}
