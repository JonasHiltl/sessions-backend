package repository

import (
	"context"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProfileRepository interface {
	Create(ctx context.Context, u datastruct.Profile) (datastruct.Profile, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, u datastruct.Profile) (datastruct.Profile, error)
	GetMany(ctx context.Context, ids []string) ([]datastruct.Profile, error)
	GetById(ctx context.Context, id string) (datastruct.Profile, error)
	GetByUsername(ctx context.Context, username string) (datastruct.Profile, error)
	UsernameTaken(ctx context.Context, username string) bool
	IncrementFriendCount(ctx context.Context, id string) error
	DecrementFriendCount(ctx context.Context, id string) error
}

type profileRepository struct {
	col *mongo.Collection
}

func (pq *profileRepository) Create(ctx context.Context, u datastruct.Profile) (datastruct.Profile, error) {
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

func (pq *profileRepository) GetById(ctx context.Context, idStr string) (res datastruct.Profile, err error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return res, errors.New("invalid profile id")
	}

	err = pq.
		col.
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (pq *profileRepository) GetMany(ctx context.Context, idsStr []string) (res []datastruct.Profile, err error) {
	var ids []primitive.ObjectID
	for _, i := range idsStr {
		id, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			return res, errors.New("invalid profile id")
		}
		ids = append(ids, id)
	}

	filter := bson.M{"_id": bson.M{"$in": ids}}

	cur, err := pq.col.Find(ctx, filter)
	if err != nil {
		return res, err
	}

	err = cur.All(ctx, &res)
	log.Println(err)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (pq *profileRepository) Delete(ctx context.Context, idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid profile id")
	}

	res, err := pq.
		col.
		DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("failed to delete user")
	}

	return nil
}

func (pq *profileRepository) Update(ctx context.Context, u datastruct.Profile) (res datastruct.Profile, err error) {
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

func (pq *profileRepository) GetByUsername(ctx context.Context, username string) (res datastruct.Profile, err error) {
	err = pq.
		col.
		FindOne(ctx, bson.M{"username": username}).
		Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (pq *profileRepository) UsernameTaken(ctx context.Context, username string) bool {
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

func (pq *profileRepository) IncrementFriendCount(ctx context.Context, idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid profile id")
	}

	update := bson.D{
		primitive.E{
			Key: "$inc",
			Value: bson.D{
				primitive.E{
					Key:   "friend_count",
					Value: 1,
				},
			},
		},
	}

	res, err := pq.col.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}
	log.Println(res)

	return nil
}

func (pq *profileRepository) DecrementFriendCount(ctx context.Context, idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid profile id")
	}

	update := bson.D{
		primitive.E{
			Key: "$inc",
			Value: bson.D{
				primitive.E{
					Key:   "friend_count",
					Value: -1,
				},
			},
		},
	}

	res, err := pq.col.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}
	log.Println(res)

	return nil
}
