package datasources

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresdbBuilder struct {
	client *sqlx.DB
}

func NewPostgresdbBuilder() *postgresdbBuilder {

	// "database-url", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	// dbURL := "postgres://propertylistingadmin:password@10.38.195.235:5432/propertylisting?sslmode=disable"

	host := "10.38.195.235"
	port := "5432"

	dsn := url.URL{
		// Scheme: "postgres",
		Scheme: "postgres",
		User:   url.UserPassword("propertylistingadmin", "password"),
		Host:   fmt.Sprintf("%s:%s", host, port),
		Path:   "propertylisting",
	}
	q := dsn.Query()
	q.Add("sslmode", "disable")
	dsn.RawQuery = q.Encode()
	dbURL := dsn.String()

	conn := sqlx.MustConnect("postgres", dbURL)

	return &postgresdbBuilder{
		client: conn,
	}
}

func (b *postgresdbBuilder) DBClient() interface{} {
	return b.client // replace this!!
}
func (b *postgresdbBuilder) migrateDb(ctx context.Context, db *sql.DB) error {
	return nil
}
func (b *postgresdbBuilder) waitForDB(ctx context.Context, conn Conn) error {
	return nil // replace this!!
}

func (b *postgresdbBuilder) Health(ctx context.Context) error {
	return nil // replace this!!
}

func (b *postgresdbBuilder) getDatabaseClient() DatabaseClient {
	return DatabaseClient{
		Client: b.client,
	}
}
