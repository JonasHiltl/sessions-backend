package repository

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PartyQuery interface {
	Create(ctx context.Context, p datastruct.Party) (datastruct.Party, error)
	Update(ctx context.Context, p datastruct.Party) (datastruct.Party, error)
	Delete(ctx context.Context, pId string) error
	GetById(ctx context.Context, pId string) (datastruct.Party, error)
}

type partyQuery struct {
	partyCol *mongo.Collection
}

func (pq *partyQuery) Create(ctx context.Context, p datastruct.Party) (datastruct.Party, error) {
	result, err := pq.partyCol.InsertOne(ctx, p)
	if err != nil {
		return datastruct.Party{}, err
	}

	p.ID = result.InsertedID.(primitive.ObjectID)

	return p, nil
}

func (pq *partyQuery) GetById(ctx context.Context, pId string) (datastruct.Party, error) {
	_id, err := primitive.ObjectIDFromHex(pId)
	if err != nil {
		return datastruct.Party{}, err
	}

	var result datastruct.Party
	err = pq.partyCol.FindOne(ctx, bson.D{{Key: "_id", Value: _id}}).Decode(&result)
	if err != nil {
		return datastruct.Party{}, nil
	}

	return result, nil
}

func (pq *partyQuery) Update(ctx context.Context, p datastruct.Party) (datastruct.Party, error) {
	filter := bson.M{"_id": p.ID}
	value := bson.M{}
	if p.Title != "" {
		value["title"] = p.Title
	}
	if len(p.Location.Coordinates) >= 0 {
		value["location"] = p.Location
	}
	if p.CreatorId != "" {
		value["creatorId"] = p.CreatorId
	}
	update := bson.D{{
		Key:   "$set",
		Value: value,
	}}

	var result datastruct.Party
	err := pq.partyCol.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&result)
	if err != nil {
		return datastruct.Party{}, err
	}

	return result, nil
}

func (pq *partyQuery) Delete(ctx context.Context, pId string) error {
	_id, err := primitive.ObjectIDFromHex(pId)
	if err != nil {
		return err
	}

	_, err = pq.partyCol.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return err
	}
	return nil
}
