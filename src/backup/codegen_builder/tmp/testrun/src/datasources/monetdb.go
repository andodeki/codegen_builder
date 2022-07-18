package datasources

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type monetdbBuilder struct {
	client *sqlx.DB
}

func newmonetdbBuilder() *monetdbBuilder {
	return &monetdbBuilder{}
}

func (b *monetdbBuilder) DBClient() interface{} {
	return b.client // replace this!!
}
func (b *monetdbBuilder) migrateDb(ctx context.Context, db *sql.DB) error {
	return nil
}
func (b *monetdbBuilder) waitForDB(ctx context.Context, conn Conn) error {
	return nil // replace this!!
}

func (b *monetdbBuilder) Health(ctx context.Context) error {
	return nil // replace this!!
}

func (b *monetdbBuilder) getDatabaseClient() DatabaseClient {
	return DatabaseClient{
		Client: b.client,
	}
}
