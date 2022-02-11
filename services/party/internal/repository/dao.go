package repository

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Dao interface {
	NewPartyQuery() PartyQuery
}

type dao struct {
	db *dynamo.DB
}

func NewDB() (*dynamo.DB, error) {
	host := os.Getenv("DYNAMO_HOST")
	port := os.Getenv("DYNAMO_PORT")
	url := fmt.Sprintf("%v:%v", host, port)

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("eu-central-1"),
		Endpoint: aws.String(url)})

	if err != nil {
		return nil, err
	}

	db := dynamo.New(sess)

	return db, nil
}

func NewDAO(db *dynamo.DB) Dao {
	return &dao{db: db}
}

func (d *dao) NewPartyQuery() PartyQuery {
	return &partyQuery{db: d.db}
}
