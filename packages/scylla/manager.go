package scylla

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Manager struct {
	keyspace string
	hosts    []string
}

func NewManager(keyspace string, hosts []string) *Manager {
	return &Manager{
		keyspace: keyspace,
		hosts:    hosts,
	}
}

func (m *Manager) Connect() (gocqlx.Session, error) {
	return m.connect(m.keyspace, m.hosts)
}

func (m *Manager) connect(keyspace string, hosts []string) (gocqlx.Session, error) {
	c := gocql.NewCluster(hosts...)
	c.Keyspace = keyspace
	return gocqlx.WrapSession(c.CreateSession())
}

func (m *Manager) CreateKeyspace(keyspace string) error {
	session, err := m.connect("system", m.hosts)
	if err != nil {
		return err
	}
	defer session.Close()

	stmt := `CREATE KEYSPACE IF NOT EXISTS sessions WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`
	return session.ExecStmt(stmt)
}
