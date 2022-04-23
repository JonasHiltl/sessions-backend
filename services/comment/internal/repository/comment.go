package repository

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/datastruct"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

const (
	TABLE_NAME        = "comments"
	COMMENTS_BY_PARTY = "comments_by_party"
)

var commentsMetadata = table.Metadata{
	Name:    TABLE_NAME,
	Columns: []string{"id", "party_id", "author_id", "body"},
	PartKey: []string{"id"},
}

type CommentRepository interface {
	Create(ctx context.Context, p datastruct.Comment) (datastruct.Comment, error)
	Delete(ctx context.Context, uId, cId string) error
	GetByParty(ctx context.Context, pId string, page []byte, limit uint32) ([]datastruct.Comment, []byte, error)
	GetByPartyUser(ctx context.Context, pId, uId string) ([]datastruct.Comment, error)
}

type commentRepository struct {
	sess *gocqlx.Session
}

func (cq *commentRepository) Create(ctx context.Context, c datastruct.Comment) (datastruct.Comment, error) {
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

// https://github.com/scylladb/scylla/issues/10171
// TODO: currently deletion by index is not supported if supported create GSI on comment_id and delete by it
func (cq *commentRepository) Delete(ctx context.Context, uId, cId string) error {
	stmt, names := qb.
		Delete(TABLE_NAME).
		Where(qb.Eq("id")).
		If(qb.Eq("author_id")).
		ToCql()

	log.Println(stmt)
	log.Println(uId)
	log.Println(cId)

	err := cq.sess.
		Query(stmt, names).
		BindMap((qb.M{"id": cId, "author_id": uId})).
		ExecRelease()
	if err != nil {
		if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
			return errors.New("you can only Delete your own Comments")
		}
		return err
	}
	return nil
}

func (cq *commentRepository) GetByParty(ctx context.Context, pId string, page []byte, limit uint32) ([]datastruct.Comment, []byte, error) {
	var result []datastruct.Comment
	stmt, names := qb.
		Select(COMMENTS_BY_PARTY).
		Where(qb.Eq("party_id")).
		ToCql()

	q := cq.sess.
		Query(stmt, names).
		BindMap((qb.M{"party_id": pId}))
	defer q.Release()

	q.PageState(page)
	if limit == 0 {
		q.PageSize(10)
	} else {
		q.PageSize(int(limit))
	}

	iter := q.Iter()
	err := iter.Select(&result)
	if err != nil {
		log.Println(err)
		return []datastruct.Comment{}, nil, errors.New("no comments found")
	}

	return result, iter.PageState(), nil
}

func (cq *commentRepository) GetByPartyUser(ctx context.Context, pId, uId string) ([]datastruct.Comment, error) {
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
