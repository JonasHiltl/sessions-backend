package repository

import (
	"errors"
	"os"
	"strings"

	"github.com/jonashiltl/sessions-backend/packages/scylla"
	"github.com/scylladb/gocqlx/v2"
)

type Dao interface {
	NewPartyQuery() PartyQuery
}

type dao struct {
	sess *gocqlx.Session
}

func NewDB() (gocqlx.Session, error) {
	keyspace, exists := os.LookupEnv("SCYLLA_KEYSPACE")
	if !exists {
		return gocqlx.Session{}, errors.New("scylla keyspace not defined")
	}
	hosts, exists := os.LookupEnv("SCYLLA_HOSTS")
	if !exists {
		return gocqlx.Session{}, errors.New("scylla hosts not defined")
	}
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

func (d *dao) NewPartyQuery() PartyQuery {
	return &partyQuery{sess: d.sess}
}
