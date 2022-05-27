package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"

	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jonashiltl/sessions-backend/services/party/repository/migrations"
	"github.com/scylladb/gocqlx/v2"
)

type Dao interface {
	NewPartyRepository() PartyRepository
}

type dao struct {
	sess *gocqlx.Session
}

func NewDB(dbUser, dbPW, dbName, dbHost string, dbPort uint16) (*pgxpool.Pool, error) {
	urlStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPW, dbHost, fmt.Sprint(dbPort), dbName)
	pgURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	connURL := *pgURL
	if connURL.Scheme == "cockroachdb" {
		connURL.Scheme = "postgres"
	}

	c, err := pgxpool.ParseConfig(connURL.String())
	if err != nil {
		return nil, err
	}
	c.ConnConfig.LogLevel = pgx.LogLevelDebug

	pool, err := pgxpool.ConnectConfig(context.Background(), c)
	if err != nil {
		return nil, fmt.Errorf("pgx connection error: %w", err)
	}

	db := stdlib.OpenDB(*c.ConnConfig)
	defer db.Close()

	err = validateSchema(db)
	if err != nil {
		log.Printf("Schema validation error: %v", err)
	}

	return pool, nil
}

func NewDAO(conn *pgxpool.Pool) Dao {
	return &dao{sess: nil}
}

func (d *dao) NewPartyRepository() PartyRepository {
	return &partyRepository{sess: d.sess}
}

const version = 1

func validateSchema(db *sql.DB) error {
	sourceInstance, err := bindata.WithInstance(bindata.Resource(migrations.AssetNames(), migrations.Asset))
	if err != nil {
		return err
	}

	targetInstance, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("go-bindata", sourceInstance, "postgres", targetInstance)
	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}
