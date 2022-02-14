package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/guregu/dynamo"
	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
)

var TABLE_NAME string = "sessions"

type PartyQuery interface {
	Create(ctx context.Context, p datastruct.Party) (datastruct.Party, error)
	Update(ctx context.Context, p datastruct.Party) (datastruct.Party, error)
	Delete(ctx context.Context, cId, pId string) error
	Get(ctx context.Context, pId string) (datastruct.Party, error)
	GeoSearch(ctx context.Context, nHashes []string) ([]datastruct.Party, error)
}

type partyQuery struct {
	db *dynamo.DB
}

func (pq *partyQuery) Create(ctx context.Context, p datastruct.Party) (datastruct.Party, error) {
	v := validator.New()
	err := v.Struct(p)
	if err != nil {
		return datastruct.Party{}, err
	}

	table := pq.db.Table(TABLE_NAME)

	err = table.Put(p).If("attribute_not_exists(geohash)").RunWithContext(ctx)
	if err != nil {
		return datastruct.Party{}, err
	}

	return p, nil
}

func (pq *partyQuery) Get(ctx context.Context, pId string) (datastruct.Party, error) {
	var result datastruct.Party

	table := pq.db.Table(TABLE_NAME)

	var sb strings.Builder
	sb.WriteString("P#")
	sb.WriteString(pId)

	err := table.
		Get("pk", sb.String()).
		Range("sk", dynamo.Equal, "party").
		Filter("'ttl' >= ?", time.Now().Unix()).
		OneWithContext(ctx, &result)
	if err != nil {
		return datastruct.Party{}, err
	}

	return result, nil
}

func (pq *partyQuery) Update(ctx context.Context, p datastruct.Party) (datastruct.Party, error) {
	table := pq.db.Table(TABLE_NAME)

	var pId strings.Builder
	pId.WriteString("P#")
	pId.WriteString(p.Id)

	var result datastruct.Party
	update := table.
		Update("pk", pId).
		Range("sk", "party")

	if p.Title != "" {
		update.Set("title", p.Title)
	}

	if p.IsPublic != "" {
		update.Set("gsi2_pk_isPublic", p.IsPublic)
	}

	if p.GHash != "" {
		update.Set("gsi2_sk_geohash", p.GHash)
	}

	err := update.
		If("'gsi1_pk_userId' <> ?", p.CreatorId).
		RunWithContext(ctx)
	if err != nil {
		return datastruct.Party{}, err
	}

	return result, nil
}

func (pq *partyQuery) Delete(ctx context.Context, cId, pId string) error {
	table := pq.db.Table(TABLE_NAME)

	var sb strings.Builder
	sb.WriteString("P#")
	sb.WriteString(pId)

	err := table.
		Delete("pk", sb.String()).
		Range("sk", "party").
		If("gsi1_pk_userId = ?", cId).
		RunWithContext(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
			return errors.New("you can only Delete your own Parties")
		}
		return err
	}
	return nil
}

func (pq *partyQuery) GeoSearch(ctx context.Context, nHashes []string) ([]datastruct.Party, error) {
	table := pq.db.Table(TABLE_NAME)
	var result []datastruct.Party

	for _, h := range nHashes {
		var partialResult []datastruct.Party

		err := table.
			Get("pk", "false").
			Range("sk", dynamo.Equal, h).
			Index("PartyGeoSearch").
			Filter("'ttl' >= ?", time.Now().Unix()).
			AllWithContext(ctx, &partialResult)
		if err != nil {
			return []datastruct.Party{}, nil
		}

		result = append(result, partialResult...)
	}

	return result, nil
}
