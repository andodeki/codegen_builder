package datasources

import (
	"context"
)

type dbdirector struct {
	Client DatabaseBuilder
}

func NewDBDirector(b DatabaseBuilder) *dbdirector {
	return &dbdirector{
		Client: b,
	}
}

func (d *dbdirector) SetDatabaseBuilder(b DatabaseBuilder) {
	d.Client = b
}

func (d *dbdirector) BuildDatabaseClient(ctx context.Context, conn Conn) interface{} {
	d.Client.migrateDb(ctx, conn.PGDBconn)
	d.Client.getDatabaseClient()
	d.Client.waitForDB(ctx, conn)
	d.Client.Health(ctx)
	return d.Client.DBClient()
}
