package repository

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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

	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("eu-central-1"),
		Endpoint: aws.String(url),
		Credentials: credentials.NewCredentials(&credentials.StaticProvider{
			Value: credentials.Value{
				AccessKeyID:     "dummy",
				SecretAccessKey: "dummy",
				SessionToken:    "dummy",
				ProviderName:    "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	}))

	db := dynamo.New(sess)

	/* 	input := &dynamodb.CreateTableInput{
	   		TableName: aws.String(TABLE_NAME),
	   		AttributeDefinitions: []*dynamodb.AttributeDefinition{
	   			{
	   				AttributeName: aws.String("pk"),
	   				AttributeType: aws.String("S"),
	   			},
	   			{
	   				AttributeName: aws.String("sk"),
	   				AttributeType: aws.String("S"),
	   			},
	   			{
	   				AttributeName: aws.String("gsi1_pk_userId"),
	   				AttributeType: aws.String("S"),
	   			},
	   			{
	   				AttributeName: aws.String("gsi2_pk_isPublic"),
	   				AttributeType: aws.String("S"),
	   			},
	   			{
	   				AttributeName: aws.String("gsi2_sk_geohash"),
	   				AttributeType: aws.String("S"),
	   			},
	   		},
	   		KeySchema: []*dynamodb.KeySchemaElement{
	   			{
	   				AttributeName: aws.String("pk"),
	   				KeyType:       aws.String("HASH"),
	   			},
	   			{
	   				AttributeName: aws.String("sk"),
	   				KeyType:       aws.String("RANGE"),
	   			},
	   		},
	   		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
	   			{
	   				IndexName: aws.String("ByUserId"),
	   				KeySchema: []*dynamodb.KeySchemaElement{
	   					{
	   						AttributeName: aws.String("gsi1_pk_userId"),
	   						KeyType:       aws.String("HASH"),
	   					},
	   					{
	   						AttributeName: aws.String("sk"),
	   						KeyType:       aws.String("RANGE"),
	   					},
	   				},
	   				Projection: &dynamodb.Projection{
	   					ProjectionType:   aws.String("INCLUDE"),
	   					NonKeyAttributes: aws.StringSlice([]string{"pk", "title", "gsi2_pk_isPublic", "geohash", "taggedFriends", "views", "content", "isViewed", "ttl"}),
	   				},
	   				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
	   					ReadCapacityUnits:  aws.Int64(1),
	   					WriteCapacityUnits: aws.Int64(1),
	   				},
	   			},
	   			{
	   				IndexName: aws.String("PartyGeoSearch"),
	   				KeySchema: []*dynamodb.KeySchemaElement{
	   					{
	   						AttributeName: aws.String("gsi2_pk_isPublic"),
	   						KeyType:       aws.String("HASH"),
	   					},
	   					{
	   						AttributeName: aws.String("gsi2_sk_geohash"),
	   						KeyType:       aws.String("RANGE"),
	   					},
	   				},
	   				Projection: &dynamodb.Projection{
	   					ProjectionType:   aws.String("INCLUDE"),
	   					NonKeyAttributes: aws.StringSlice([]string{"pk", "gsi1_pk_userId", "title", "ttl"}),
	   				},
	   				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
	   					ReadCapacityUnits:  aws.Int64(1),
	   					WriteCapacityUnits: aws.Int64(1),
	   				},
	   			},
	   		},
	   		BillingMode: aws.String("PROVISIONED"),
	   		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
	   			ReadCapacityUnits:  aws.Int64(1),
	   			WriteCapacityUnits: aws.Int64(1),
	   		},
	   	}

	   	_, err := db.Client().CreateTable(input)
	   	if err != nil {
	   		return nil, err
	   	}
	*/
	return db, nil
}

func NewDAO(db *dynamo.DB) Dao {
	return &dao{db: db}
}

func (d *dao) NewPartyQuery() PartyQuery {
	return &partyQuery{db: d.db}
}
