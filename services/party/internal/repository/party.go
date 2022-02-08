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
	Search(ctx context.Context, q string, page int) ([]datastruct.Party, error)
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
	if p.Location.Coordinates[0] != 0 && p.Location.Coordinates[1] != 0 {
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

func (pq *partyQuery) Search(ctx context.Context, q string, page int) ([]datastruct.Party, error) {
	var matches []datastruct.Party
	var perPage int64 = 10

	if page <= 0 {
		page = 1
	}

	filter := bson.D{
		{
			Key: "$text",
			Value: bson.D{{
				Key:   "$search",
				Value: q,
			}},
		},
		{
			Key:   "isGlobal",
			Value: true,
		},
	}
	sort := bson.D{{
		Key: "score",
		Value: bson.D{{
			Key:   "$meta",
			Value: "textScore",
		}},
	}}
	options := options.Find().SetSort(sort)
	options.SetSkip((int64(page) - 1) * perPage)
	options.SetLimit(perPage)

	cursor, err := pq.partyCol.Find(
		ctx,
		filter,
		options,
	)
	if err != nil {
		return []datastruct.Party{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var party datastruct.Party
		cursor.Decode(&party)

		matches = append(matches, party)
	}

	// return empty array instead of null when parsing json
	if len(matches) == 0 {
		matches = make([]datastruct.Party, 0)
	}

	return matches, nil
}
