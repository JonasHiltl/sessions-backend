package repository

import (
	"context"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/guregu/dynamo"
	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
)

var TABLE_NAME string = "sessions"

type PartyQuery interface {
	Create(ctx context.Context, p datastruct.Party) (datastruct.Party, error)
	Update(ctx context.Context, cId, pId, title string) (datastruct.Party, error)
	Delete(ctx context.Context, cId, pId string) error
	Get(ctx context.Context, cId, pId string) (datastruct.Party, error)
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

func (pq *partyQuery) Get(ctx context.Context, cId string, pId string) (datastruct.Party, error) {
	var result datastruct.Party

	table := pq.db.Table(TABLE_NAME)

	var sb strings.Builder
	sb.WriteString("PARTY#")
	sb.WriteString(pId)

	err := table.Get("pk", cId).
		Range("sk", dynamo.Equal, sb.String()).
		OneWithContext(ctx, result)
	if err != nil {
		return datastruct.Party{}, nil
	}

	return result, nil
}

func (pq *partyQuery) Update(ctx context.Context, cId string, pId string, title string) (datastruct.Party, error) {
	table := pq.db.Table(TABLE_NAME)

	var sb strings.Builder
	sb.WriteString("PARTY#")
	sb.WriteString(pId)

	var result datastruct.Party
	update := table.Update("pk", cId).Range("sk", sb.String())

	if title != "" {
		update.Set("title", title)
	}

	err := update.RunWithContext(ctx)
	if err != nil {
		return datastruct.Party{}, err
	}

	return result, nil
}

func (pq *partyQuery) Delete(ctx context.Context, cId string, pId string) error {
	table := pq.db.Table(TABLE_NAME)

	var sb strings.Builder
	sb.WriteString("PARTY#")
	sb.WriteString(pId)

	err := table.Delete("pk", cId).
		Range("sk", sb.String()).RunWithContext(ctx)
	if err != nil {
		return err
	}
	return nil
}
