package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Dao interface {
	NewPartyQuery() PartyQuery
}

type dao struct {
	partyCol *mongo.Collection
}

func NewDB() (*mongo.Collection, error) {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	url := fmt.Sprintf("mongodb://%v:%v", host, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database("sessions")
	partyCol := db.Collection(datastruct.PartyCollectionName)

	return partyCol, nil
}

func NewDAO(partyCol *mongo.Collection) Dao {
	return &dao{partyCol: partyCol}
}

func (d *dao) NewPartyQuery() PartyQuery {
	return &partyQuery{partyCol: d.partyCol}
}
