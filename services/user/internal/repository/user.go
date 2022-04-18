package repository

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Create(context.Context, datastruct.User) (datastruct.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, u datastruct.User) (datastruct.User, error)
	UpdateVerified(ctx context.Context, email string, emailVerified bool) (datastruct.User, error)
	GetById(ctx context.Context, id string) (datastruct.User, error)
	RotateEmailCode(ctx context.Context, email string) (datastruct.User, error)
	EmailTaken(ctx context.Context, email string) bool
	UsernameTaken(ctx context.Context, username string) bool
	GetByEmail(ctx context.Context, email string) (datastruct.User, error)
	GetByEmailOrUsername(ctx context.Context, usernameOrEmail string) (datastruct.User, error)
}

type userRepository struct {
	col *mongo.Collection
}

func (r *userRepository) Create(ctx context.Context, u datastruct.User) (datastruct.User, error) {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return datastruct.User{}, err
	}

	code, err := utils.GenerateOTP(4)
	if err != nil {
		return datastruct.User{}, errors.New("no user found")
	}

	u.EmailCode = code

	res, err := r.
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

func (uq *userRepository) Delete(ctx context.Context, idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid profile id")
	}

	res, err := uq.
		col.
		DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("failed to delete user entry of user")
	}

	return nil
}

func (uq *userRepository) Update(ctx context.Context, u datastruct.User) (res datastruct.User, err error) {
	input := bson.M{}
	filter := bson.M{"_id": u.Id}

	after := options.After
	opt := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	if u.Provider != "" {
		input["provider"] = u.Provider
	}

	if u.Email != "" {
		input["email"] = u.Email
	}

	if u.PasswordHash != "" {
		input["password"] = u.PasswordHash
	}

	if u.Role != "" {
		input["role"] = u.Role
	}

	if u.PasswordHash != "" {
		input["password_hash"] = u.PasswordHash
	}

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
		FindOneAndUpdate(ctx, filter, bson.M{"$set": input}, &opt).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *userRepository) GetById(ctx context.Context, idStr string) (res datastruct.User, err error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return res, errors.New("invalid profile id")
	}

	err = r.
		col.
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&res)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return res, errors.New("no user found")
		}
		return res, err
	}

	return res, err
}

func (r *userRepository) RotateEmailCode(ctx context.Context, email string) (res datastruct.User, err error) {
	code, err := utils.GenerateOTP(4)
	if err != nil {
		return res, errors.New("no user found")
	}

	input := bson.M{
		"email_code": code,
	}
	filter := bson.M{"email": email}

	after := options.After
	opt := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	err = r.
		col.
		FindOneAndUpdate(ctx, filter, input, &opt).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *userRepository) UpdateVerified(ctx context.Context, email string, emailVerified bool) (res datastruct.User, err error) {
	input := bson.M{
		"email_verified": emailVerified,
	}
	filter := bson.M{"email": email}

	after := options.After
	opt := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	err = r.
		col.
		FindOneAndUpdate(ctx, filter, input, &opt).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *userRepository) EmailTaken(ctx context.Context, email string) bool {
	user := datastruct.User{}
	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return false
	}

	if user.Id.Hex() == "" {
		return false
	}

	return true

}

func (r *userRepository) UsernameTaken(ctx context.Context, username string) bool {
	user := datastruct.Profile{}
	err := r.col.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return false
	}

	if user.Id.Hex() == "" {
		return false
	}

	return true
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (res datastruct.User, err error) {
	err = r.
		col.
		FindOne(ctx, bson.M{
			"email": email,
		},
		).
		Decode(&res)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return res, errors.New("no user found")
		}
		return res, err
	}

	return res, err
}

func (r *userRepository) GetByEmailOrUsername(ctx context.Context, usernameOrEmail string) (res datastruct.User, err error) {
	err = r.
		col.
		FindOne(ctx, bson.M{
			"$or": []interface{}{
				bson.M{"username": usernameOrEmail},
				bson.M{"email": usernameOrEmail},
			},
		},
		).
		Decode(&res)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return res, errors.New("no user found")
		}
		return res, err
	}

	return res, err
}
