package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

const TABLE_NAME string = "party"

var partyMetadata = table.Metadata{
	Name:    TABLE_NAME,
	Columns: []string{"id", "user_id", "title", "is_public", "geohash", "created_at"},
	PartKey: []string{"id"},
}
var partyTable = table.New(partyMetadata)

type PartyQuery interface {
	Create(ctx context.Context, p datastruct.Party, ttl time.Duration) (datastruct.Party, error)
	Update(ctx context.Context, p datastruct.Party) error
	Delete(ctx context.Context, uId, pId string) error
	Get(ctx context.Context, pId string) (datastruct.Party, error)
	GetByUser(ctx context.Context, uId string) ([]datastruct.Party, error)
	GeoSearch(ctx context.Context, nHashes []string) ([]datastruct.Party, error)
}

type partyQuery struct {
	sess *gocqlx.Session
}

func (pq *partyQuery) Create(ctx context.Context, p datastruct.Party, ttl time.Duration) (datastruct.Party, error) {
	v := validator.New()
	err := v.Struct(p)
	if err != nil {
		return datastruct.Party{}, err
	}

	stmt, names := qb.
		Insert(TABLE_NAME).
		Columns(partyMetadata.Columns...).
		TTL(ttl).
		ToCql()

	err = pq.sess.
		Query(stmt, names).
		BindStruct(p).
		ExecRelease()
	if err != nil {
		return datastruct.Party{}, err
	}

	return p, nil
}

func (pq *partyQuery) Get(ctx context.Context, pId string) (datastruct.Party, error) {
	var result datastruct.Party
	err := pq.sess.
		Query(partyTable.Get()).
		BindMap((qb.M{"id": pId})).
		GetRelease(&result)
	if err != nil {
		return datastruct.Party{}, err
	}

	return result, nil
}

func (pq *partyQuery) Update(ctx context.Context, p datastruct.Party) error {
	b := qb.
		Update(TABLE_NAME).
		Where(qb.Eq("id"))

	if p.Title != "" {
		b.Set("title")
	}

	if p.GHash != "" {
		b.Set("geohash")
	}

	b.If(qb.Eq("user_id"))
	stmt, names := b.ToCql()

	err := pq.sess.Query(stmt, names).
		BindMap((qb.M{"id": p.Id, "title": p.Title, "geohash": p.GHash, "user_id": p.UId})).
		ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

func (pq *partyQuery) Delete(ctx context.Context, uId, pId string) error {
	stmt, names := qb.
		Delete(TABLE_NAME).
		Where(qb.Eq("id")).
		If(qb.Eq("user_id")).
		ToCql()

	err := pq.sess.
		Query(stmt, names).
		BindMap((qb.M{"id": pId, "user_id": uId})).
		ExecRelease()
	if err != nil {
		if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
			return errors.New("you can only Delete your own Parties")
		}
		return err
	}
	return nil
}

func (pq *partyQuery) GetByUser(ctx context.Context, uId string) ([]datastruct.Party, error) {
	var result []datastruct.Party
	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("user_id")).
		OrderBy("created_at", qb.DESC).
		ToCql()

	err := pq.sess.
		Query(stmt, names).
		BindMap((qb.M{"user_id": uId})).
		GetRelease(&result)
	if err != nil {
		return []datastruct.Party{}, err
	}

	return result, nil
}

func (pq *partyQuery) GeoSearch(ctx context.Context, nHashes []string) ([]datastruct.Party, error) {
	var result []datastruct.Party

	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("is_public")).
		Where(qb.In("geohash")).
		OrderBy("created_at", qb.DESC).
		ToCql()

	iter := pq.sess.
		Query(stmt, names).
		BindMap((qb.M{"is_public": false, "geohash": nHashes})).
		PageSize(10).
		Iter()

	for iter.Scan() {

	}

	return result, nil
}
