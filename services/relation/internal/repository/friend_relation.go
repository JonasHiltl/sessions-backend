package repository

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/datastruct"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

const (
	TABLE_NAME                string = "friend_relations"
	FILTERED_FRIEND_RELATIONS string = "filtered_friend_relations"
)

var friendRelationMetadata = table.Metadata{
	Name:    TABLE_NAME,
	Columns: []string{"user_id", "friend_id", "accepted", "created_at"},
	PartKey: []string{"user_id", "friend_id"},
}
var friendRelationTable = table.New(friendRelationMetadata)

type FriendRelationRepository interface {
	CreateFriendRelation(ctx context.Context, fr datastruct.FriendRelation) error
	AcceptFriendRelation(ctx context.Context, uId, fId string) error
	GetFriendRelation(ctx context.Context, uId, fId string) (datastruct.FriendRelation, error)
	GetFriendsOfUser(ctx context.Context, uId string, page []byte, limit uint32) ([]datastruct.FriendRelation, []byte, error)
}

type friendRelationRepository struct {
	sess *gocqlx.Session
}

func (r *friendRelationRepository) CreateFriendRelation(ctx context.Context, fr datastruct.FriendRelation) error {
	v := validator.New()
	err := v.Struct(fr)
	if err != nil {
		return err
	}

	stmt, names := qb.
		Insert(TABLE_NAME).
		Unique().
		Columns(friendRelationMetadata.Columns...).
		ToCql()

	err = r.sess.
		Query(stmt, names).
		BindStruct(fr).
		ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

func (r *friendRelationRepository) AcceptFriendRelation(ctx context.Context, uId, fId string) error {
	stmt, names := qb.
		Update(TABLE_NAME).
		Where(qb.Eq("user_id")).
		Where(qb.Eq("friend_id")).
		If(qb.EqLit("accepted", "false")).
		Set("accepted").
		Set("accepted_at").
		ToCql()

	err := r.sess.Query(stmt, names).
		BindMap((qb.M{
			"user_id":     uId,
			"friend_id":   fId,
			"accepted":    true,
			"accepted_at": time.Now(),
		})).
		ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

func (r *friendRelationRepository) GetFriendRelation(ctx context.Context, uId, fId string) (res datastruct.FriendRelation, err error) {
	err = r.sess.
		Query(friendRelationTable.Get()).
		BindMap((qb.M{"user_id": uId, "friend_id": fId})).
		GetRelease(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *friendRelationRepository) GetFriendsOfUser(ctx context.Context, uId string, page []byte, limit uint32) (result []datastruct.FriendRelation, nextPage []byte, err error) {
	stmt, names := qb.
		Select(FILTERED_FRIEND_RELATIONS).
		Where(qb.Eq("user_id")).
		Where(qb.Eq("accepted")).
		ToCql()

	q := r.sess.
		Query(stmt, names).
		BindMap((qb.M{
			"user_id":  uId,
			"accepted": true,
		}))
	defer q.Release()

	q.PageState(page)
	if limit == 0 {
		q.PageSize(20)
	} else {
		q.PageSize(int(limit))
	}

	iter := q.Iter()
	err = iter.Select(&result)
	if err != nil {
		return []datastruct.FriendRelation{}, nil, errors.New("no friends found")
	}

	return result, iter.PageState(), nil
}
