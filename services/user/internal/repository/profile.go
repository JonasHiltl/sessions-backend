package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileRepository interface {
	GetMany(ctx context.Context, ids []string) ([]datastruct.Profile, error)
	GetById(ctx context.Context, id string) (datastruct.Profile, error)
	GetByUsername(ctx context.Context, username string) (datastruct.User, error)
	IncrementFriendCount(ctx context.Context, id string) error
	DecrementFriendCount(ctx context.Context, id string) error
}

type profileRepository struct {
	col *mongo.Collection
}

func (r *profileRepository) GetById(ctx context.Context, idStr string) (res datastruct.Profile, err error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return res, errors.New("invalid user id")
	}

	err = r.
		col.
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&res)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return res, errors.New("no profile found")
		}
		return res, err
	}

	return res, err
}

func (r *profileRepository) GetMany(ctx context.Context, idsStr []string) (res []datastruct.Profile, err error) {
	var ids []primitive.ObjectID
	for _, i := range idsStr {
		id, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			return res, errors.New("invalid user id")
		}
		ids = append(ids, id)
	}

	filter := bson.M{"_id": bson.M{"$in": ids}}

	cur, err := r.col.Find(ctx, filter)
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

func (r *profileRepository) GetByUsername(ctx context.Context, username string) (res datastruct.User, err error) {
	err = r.
		col.
		FindOne(ctx, bson.M{
			"username": username,
		},
		).
		Decode(&res)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return res, errors.New("no profile found")
		}
		return res, err
	}

	return res, err
}

func (r *profileRepository) IncrementFriendCount(ctx context.Context, idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid user id")
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

	res, err := r.col.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}
	log.Println(res)

	return nil
}

func (r *profileRepository) DecrementFriendCount(ctx context.Context, idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid user id")
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

	res, err := r.col.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}
	log.Println(res)

	return nil
}
