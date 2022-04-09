package repository

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

const (
	TABLE_NAME    string = "party"
	PARTY_BY_USER string = "party_by_user"
)

var partyMetadata = table.Metadata{
	Name:    TABLE_NAME,
	Columns: []string{"id", "user_id", "title", "is_public", "geohash"},
	PartKey: []string{"id"},
}
var partyTable = table.New(partyMetadata)

type PartyQuery interface {
	Create(ctx context.Context, p datastruct.Party, ttl time.Duration) (datastruct.Party, error)
	Update(ctx context.Context, p datastruct.Party) error
	Delete(ctx context.Context, uId, pId string) error
	Get(ctx context.Context, pId string) (datastruct.Party, error)
	GetByUser(ctx context.Context, uId string, page []byte) ([]datastruct.Party, []byte, error)
	GeoSearch(ctx context.Context, nHashes []string, page []byte) ([]datastruct.Party, []byte, error)
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

	if p.Position[0] != 0 && p.Position[1] != 0 {
		b.Set("position")
	}

	if p.StartDate != "" {
		b.Set("start_date")
	}

	if p.StreetAddress != "" {
		b.Set("street_address")
	}

	if p.PostalCode != "" {
		b.Set("postal_code")
	}

	if p.State != "" {
		b.Set("state")
	}

	if p.Country != "" {
		b.Set("country")
	}

	b.If(qb.Eq("user_id"))
	stmt, names := b.ToCql()

	err := pq.sess.Query(stmt, names).
		BindMap((qb.M{
			"user_id":        p.UserId,
			"id":             p.Id,
			"title":          p.Title,
			"geohash":        p.GHash,
			"position":       p.Position,
			"start_date":     p.StartDate,
			"street_address": p.StreetAddress,
			"postal_code":    p.PostalCode,
			"state":          p.State,
			"country":        p.Country,
		})).
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

func (pq *partyQuery) GetByUser(ctx context.Context, uId string, page []byte) (result []datastruct.Party, nextPage []byte, err error) {
	stmt, names := qb.
		Select(PARTY_BY_USER).
		Where(qb.Eq("user_id")).
		ToCql()

	q := pq.sess.
		Query(stmt, names).
		BindMap((qb.M{"user_id": uId}))
	defer q.Release()

	q.PageState(page)
	q.PageSize(10)

	iter := q.Iter()
	err = iter.Select(&result)
	if err != nil {
		return []datastruct.Party{}, nil, errors.New("no parties found")
	}

	return result, iter.PageState(), nil
}

func (pq *partyQuery) GeoSearch(ctx context.Context, nHashes []string, page []byte) (result []datastruct.Party, nextPage []byte, err error) {
	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("is_public")).
		Where(qb.In("geohash")).
		OrderBy("created_at", qb.DESC).
		ToCql()

	q := pq.sess.
		Query(stmt, names).
		BindMap((qb.M{"is_public": false, "geohash": nHashes}))

	if len(page) > 0 {
		q.PageState(page)
	}

	q.PageSize(10)
	iter := q.Iter()

	err = iter.Select(&result)
	if err != nil {
		log.Println(err)
		return []datastruct.Party{}, nil, errors.New("no parties found")
	}

	return result, iter.PageState(), err
}
