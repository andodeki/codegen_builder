package datasources

import (
	"context"
	"database/sql"
)

type DatabaseBuilder interface {
	DBClient() interface{}
	waitForDB(ctx context.Context, conn Conn) error
	Health(ctx context.Context) error
	getDatabaseClient() DatabaseClient
	migrateDb(ctx context.Context, db *sql.DB) error
}

func GetDatabaseBuilder(builderType string) DatabaseBuilder {
	if builderType == "postgresdb" {
		db := NewPostgresdbBuilder()
		return &postgresdbBuilder{
			client: db.client,
		}
	}
	if builderType == "scylladb" {
		return &scylladbBuilder{}
	}
	if builderType == "redisdb" {
		return &redisdbBuilder{}
	}
	if builderType == "monetdb" {
		return &monetdbBuilder{}
	}
	return nil
}
