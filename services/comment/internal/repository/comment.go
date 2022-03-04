package repository

import (
	"context"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/datastruct"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

const (
	TABLE_NAME        = "comments"
	COMMENTS_BY_PARTY = "commenty_by_party"
)

var commentsMetadata = table.Metadata{
	Name:    TABLE_NAME,
	Columns: []string{"comment_id", "party_id", "author_id", "body", "created_at"},
	PartKey: []string{"party_id"},
	SortKey: []string{"author_id", "created_at"},
}

type CommentQuery interface {
	Create(ctx context.Context, p datastruct.Comment) (datastruct.Comment, error)
	Delete(ctx context.Context, uId, pId, cId string) error
	GetByParty(ctx context.Context, pId string) ([]datastruct.Comment, error)
	GetByPartyUser(ctx context.Context, pId, uId string) ([]datastruct.Comment, error)
}

type commentQuery struct {
	sess *gocqlx.Session
}

func (cq *commentQuery) Create(ctx context.Context, c datastruct.Comment) (datastruct.Comment, error) {
	v := validator.New()
	err := v.Struct(c)
	if err != nil {
		return datastruct.Comment{}, err
	}

	stmt, names := qb.
		Insert(TABLE_NAME).
		Columns(commentsMetadata.Columns...).
		ToCql()

	err = cq.sess.
		Query(stmt, names).
		BindStruct(c).
		ExecRelease()
	if err != nil {
		return datastruct.Comment{}, err
	}

	return c, err
}

// TODO: currently deletion by index is not supported if supported create GSI on comment_id and delete by it
func (cq *commentQuery) Delete(ctx context.Context, uId, pId, cId string) error {
	stmt, names := qb.
		Delete(TABLE_NAME).
		Where(qb.Eq("comment_id")).
		Where(qb.Eq("party_id")).
		If(qb.Eq("author_id")).
		ToCql()

	err := cq.sess.
		Query(stmt, names).
		BindMap((qb.M{"comment_id": cId, "party_id": pId, "author_id": uId})).
		ExecRelease()
	if err != nil {
		return err
	}
	return nil
}

func (cq *commentQuery) GetByParty(ctx context.Context, pId string) ([]datastruct.Comment, error) {
	var result []datastruct.Comment
	stmt, names := qb.
		Select(COMMENTS_BY_PARTY).
		Where(qb.Eq("party_id")).
		OrderBy("created_at", qb.DESC).
		ToCql()

	err := cq.sess.
		Query(stmt, names).
		BindMap((qb.M{"party_id": pId})).
		PageSize(10).
		Iter().
		Select(&result)
	if err != nil {
		return []datastruct.Comment{}, errors.New("no comments found")
	}

	return result, nil
}

func (cq *commentQuery) GetByPartyUser(ctx context.Context, pId, uId string) ([]datastruct.Comment, error) {
	var result []datastruct.Comment
	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("party_id")).
		Where(qb.Eq("author_id")).
		OrderBy("created_at", qb.ASC).
		ToCql()

	log.Println(stmt)

	err := cq.sess.
		Query(stmt, names).
		BindMap((qb.M{"party_id": pId, "author_id": uId})).
		PageSize(10).
		Iter().
		Select(&result)
	if err != nil {
		log.Println(err)
		return []datastruct.Comment{}, errors.New("no comments found")
	}

	return result, nil
}
