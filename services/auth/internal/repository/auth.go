package repository

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthRepository interface {
	Create(context.Context, datastruct.AuthUser) (datastruct.AuthUser, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, u datastruct.AuthUser) (datastruct.AuthUser, error)
	GetById(ctx context.Context, id string) (datastruct.AuthUser, error)
	GetByEmail(ctx context.Context, email string) (datastruct.AuthUser, error)
	RotateEmailCode(ctx context.Context, email string) (datastruct.AuthUser, error)
	UpdateVerified(ctx context.Context, email string, emailVerified bool) (datastruct.AuthUser, error)
	EmailTaken(ctx context.Context, email string) bool
}

type authRepository struct {
	col *mongo.Collection
}

func (r *authRepository) Create(ctx context.Context, u datastruct.AuthUser) (datastruct.AuthUser, error) {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return datastruct.AuthUser{}, err
	}

	code, err := utils.GenerateOTP(4)
	if err != nil {
		return datastruct.AuthUser{}, errors.New("No user found")
	}

	u.EmailCode = code

	res, err := r.
		col.
		InsertOne(ctx, u)
	if err != nil {
		return datastruct.AuthUser{}, err
	}

	id := res.InsertedID.(primitive.ObjectID)

	if res.InsertedID != nil {
		u.Id = id
	}

	return u, nil
}

func (uq *authRepository) Delete(ctx context.Context, idStr string) error {
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
		return errors.New("Failed to delete auth entry of user")
	}

	return nil
}

func (uq *authRepository) Update(ctx context.Context, u datastruct.AuthUser) (res datastruct.AuthUser, err error) {
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

	err = uq.
		col.
		FindOneAndUpdate(ctx, filter, input, &opt).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *authRepository) GetById(ctx context.Context, idStr string) (res datastruct.AuthUser, err error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return res, errors.New("invalid profile id")
	}

	err = r.
		col.
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (r *authRepository) GetByEmail(ctx context.Context, email string) (res datastruct.AuthUser, err error) {
	err = r.
		col.
		FindOne(ctx, bson.M{"email": email}).
		Decode(&res)
	if err != nil {
		return res, errors.New("No user found")
	}

	return res, err
}

func (r *authRepository) RotateEmailCode(ctx context.Context, email string) (res datastruct.AuthUser, err error) {
	code, err := utils.GenerateOTP(4)
	if err != nil {
		return res, errors.New("No user found")
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

func (r *authRepository) UpdateVerified(ctx context.Context, email string, emailVerified bool) (res datastruct.AuthUser, err error) {
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

func (r *authRepository) EmailTaken(ctx context.Context, email string) bool {
	user := datastruct.AuthUser{}
	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return false
	}

	if user.Id.Hex() == "" {
		return false
	}

	return true

}
