package datasources

import (
	"context"
	"database/sql"

	"github.com/go-redis/cache/v8"
)

type redisdbBuilder struct {
	client *cache.Cache
}

func newredisdbBuilder() *redisdbBuilder {
	return &redisdbBuilder{}
}

func (b *redisdbBuilder) DBClient() interface{} {
	return b.client // replace this!!
}
func (b *redisdbBuilder) migrateDb(ctx context.Context, db *sql.DB) error {
	return nil
}
func (b *redisdbBuilder) waitForDB(ctx context.Context, conn Conn) error {
	return nil // replace this!!
}

func (b *redisdbBuilder) Health(ctx context.Context) error {
	return nil // replace this!!
}

func (b *redisdbBuilder) getDatabaseClient() DatabaseClient {
	return DatabaseClient{
		Client: b.client,
	}
}
