package datasources

import (
	"context"
	"database/sql"

	"github.com/scylladb/gocqlx/v2"
)

type scylladbBuilder struct {
	client *gocqlx.Session
}

func newscylladbBuilder() *scylladbBuilder {
	return &scylladbBuilder{}
}

func (b *scylladbBuilder) DBClient() interface{} {
	return b.client // replace this!!
}
func (b *scylladbBuilder) migrateDb(ctx context.Context, db *sql.DB) error {
	return nil
}
func (b *scylladbBuilder) waitForDB(ctx context.Context, conn Conn) error {
	return nil // replace this!!
}

func (b *scylladbBuilder) Health(ctx context.Context) error {
	return nil // replace this!!
}

func (b *scylladbBuilder) getDatabaseClient() DatabaseClient {
	return DatabaseClient{
		Client: b.client,
	}
}
