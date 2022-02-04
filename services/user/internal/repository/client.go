package repository

import (
	"context"
	"log"
	"os"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jonashiltl/sessions-backend/services/user/ent"
	"github.com/jonashiltl/sessions-backend/services/user/ent/migrate"
)

func NewClient() (*ent.UserClient, error) {
	drv, err := sql.Open("mysql", os.Getenv("MYSQL"))
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	client := ent.NewClient(ent.Driver(drv))

	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client.User, nil
}
