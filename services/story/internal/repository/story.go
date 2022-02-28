package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/story/internal/datastruct"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

const TABLE_NAME string = "story"

var storyMetadata = table.Metadata{
	Name:    TABLE_NAME,
	Columns: []string{"id", "party_id", "user_id", "geohash", "url", "tagged_friends", "created_at"},
	PartKey: []string{"id"},
}
var storyTable = table.New(storyMetadata)

type StoryQuery interface {
	Create(c context.Context, s datastruct.Story) (datastruct.Story, error)
	Delete(c context.Context, uId, sId string) error
	Get(c context.Context, sId string) (datastruct.Story, error)
	GetByUser(c context.Context, uId string) ([]datastruct.Story, error)
	GetByParty(c context.Context, pId string) ([]datastruct.Story, error)
}

type storyQuery struct {
	sess *gocqlx.Session
}

func (sq *storyQuery) Create(c context.Context, s datastruct.Story) (datastruct.Story, error) {
	v := validator.New()
	err := v.Struct(s)
	if err != nil {
		return datastruct.Story{}, err
	}

	stmt, names := qb.
		Insert(TABLE_NAME).
		Columns(storyMetadata.Columns...).
		TTL(time.Hour * 24).
		ToCql()

	err = sq.sess.
		Query(stmt, names).
		BindStruct(s).
		ExecRelease()
	if err != nil {
		return datastruct.Story{}, err
	}

	return s, err
}

func (sq *storyQuery) Delete(c context.Context, uId, sId string) error {
	stmt, names := qb.
		Delete(TABLE_NAME).
		Where(qb.Eq("id")).
		If(qb.Eq("user_id")).
		ToCql()

	err := sq.sess.
		Query(stmt, names).
		BindMap((qb.M{"id": sId, "user_id": uId})).
		ExecRelease()
	if err != nil {
		return err
	}
	return nil
}

func (sq *storyQuery) Get(c context.Context, sId string) (datastruct.Story, error) {
	var result datastruct.Story
	err := sq.sess.
		Query(storyTable.Get()).
		BindMap((qb.M{"id": sId})).
		GetRelease(&result)
	if err != nil {
		if err.Error() == "not found" {
			return datastruct.Story{}, errors.New("story not found")
		}
		return datastruct.Story{}, err
	}

	return result, nil
}

func (sq *storyQuery) GetByUser(c context.Context, uId string) ([]datastruct.Story, error) {
	var result []datastruct.Story
	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("user_id")).
		ToCql()

	log.Println(stmt)

	err := sq.sess.
		Query(stmt, names).
		BindMap((qb.M{"user_id": uId})).
		PageSize(10).
		Iter().
		Select(&result)
	if err != nil {
		return []datastruct.Story{}, errors.New("no stories found")
	}

	return result, nil
}

func (sq *storyQuery) GetByParty(c context.Context, pId string) ([]datastruct.Story, error) {
	var result []datastruct.Story
	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("party_id")).
		ToCql()

	err := sq.sess.
		Query(stmt, names).
		BindMap((qb.M{"party_id": pId})).
		PageSize(10).
		Iter().
		Select(&result)
	if err != nil {
		return []datastruct.Story{}, errors.New("no stories found")
	}

	return result, nil
}
