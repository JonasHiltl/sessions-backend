package repository

import (
	"strings"

	"github.com/jonashiltl/sessions-backend/packages/scylla"
	"github.com/scylladb/gocqlx/v2"
)

type Dao interface {
	NewFriendRelationRepository() FriendRelationRepository
}

type dao struct {
	sess *gocqlx.Session
}

func NewDB(keyspace, hosts string) (gocqlx.Session, error) {
	h := strings.Split(hosts, ",")

	manager := scylla.NewManager(keyspace, h)

	if err := manager.CreateKeyspace(keyspace); err != nil {
		return gocqlx.Session{}, err
	}

	session, err := manager.Connect()
	if err != nil {
		return gocqlx.Session{}, err
	}
	return session, nil
}

func NewDAO(sess *gocqlx.Session) Dao {
	return &dao{sess: sess}
}

func (d *dao) NewFriendRelationRepository() FriendRelationRepository {
	return &friendRelationRepository{sess: d.sess}
}
