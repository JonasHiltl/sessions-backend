package repository

import (
	"context"
	"log"

	"github.com/jonashiltl/sessions-backend/services/user/datastruct"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

var profileMetadata = table.Metadata{
	Name:    TABLE_NAME,
	Columns: []string{"id", "username", "email", "firstname", "lastname", "avatar", "friend_count"},
	PartKey: []string{"id"},
}
var profileTable = table.New(profileMetadata)

type ProfileRepository interface {
	GetMany(ctx context.Context, ids []string) ([]datastruct.Profile, error)
	GetById(ctx context.Context, id string) (datastruct.Profile, error)
}

type profileRepository struct {
	sess *gocqlx.Session
}

func (r *profileRepository) GetById(ctx context.Context, id string) (res datastruct.Profile, err error) {
	err = r.sess.
		Query(profileTable.Get(profileMetadata.Columns...)).
		BindMap((qb.M{"id": id})).
		GetRelease(&res)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}

func (r *profileRepository) GetMany(ctx context.Context, ids []string) (res []datastruct.Profile, err error) {
	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.In("id")).
		ToCql()

	err = r.sess.Query(stmt, names).
		BindMap((qb.M{
			"id": ids,
		})).
		GetRelease(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
