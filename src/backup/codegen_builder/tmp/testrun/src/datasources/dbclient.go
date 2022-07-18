package datasources

import (
	"database/sql"

	"github.com/go-redis/cache/v8"
	"github.com/jmoiron/sqlx"
	"github.com/scylladb/gocqlx/v2"
	// "github.com/namsral/flag"
	// "github.com/jackc/pgx/v4"
	// "github.com/jackc/pgx/v4"
)

type DatabaseClient struct {
	Client interface{}
}

type DBS struct {
	PGDB *sqlx.DB
	MDB  *sqlx.DB
	SDB  *gocqlx.Session
	RDB  *cache.Cache
}

type Conn struct {
	PGDBconn *sql.DB
	SDBconn  *sql.DB
	RDBconn  *sql.DB
	MDBconn  *sql.DB
}
