package repository

import (
	"strings"

	"github.com/jonashiltl/sessions-backend/packages/cqlx"
	"github.com/scylladb/gocqlx/v2"
)

type Dao interface {
	NewUserRepository() UserRepository
	NewProfileRepository() ProfileRepository
}

type dao struct {
	sess *gocqlx.Session
}

const (
	USER_COLLECTION = "user"
)

func NewDB(keyspace, hosts string) (*gocqlx.Session, error) {
	h := strings.Split(hosts, ",")

	manager := cqlx.NewManager(keyspace, h)

	if err := manager.CreateKeyspace(keyspace); err != nil {
		return nil, err
	}

	session, err := manager.Connect()
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func NewDAO(sess *gocqlx.Session) Dao {
	return &dao{sess: sess}
}

func (d *dao) NewUserRepository() UserRepository {
	return &userRepository{sess: d.sess}
}

func (d *dao) NewProfileRepository() ProfileRepository {
	return &profileRepository{sess: d.sess}
}
