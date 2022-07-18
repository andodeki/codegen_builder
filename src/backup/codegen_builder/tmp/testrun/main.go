package main

import (
	"context"

	"github.com/andodeki/gen_project/codegen_builder/tmp/testrun/src/datasources"
)

func main() {
	conn := &datasources.Conn{}
	ctx := context.Background()

	postgresdbBuilder := datasources.GetDatabaseBuilder("postgresdb")
	director := datasources.NewDBDirector(postgresdbBuilder)
	postgresdbDatabaseClient := director.BuildDatabaseClient(ctx, *conn)
	_ = postgresdbDatabaseClient

	// postgresdbDatabaseClient.Client.PGDB
	// director.Client.DBClient()

	scylladbBuilder := datasources.GetDatabaseBuilder("scylladb")
	director.SetDatabaseBuilder(scylladbBuilder)
	scylladbDatabaseClient := director.BuildDatabaseClient(ctx, *conn)
	_ = scylladbDatabaseClient
	// fmt.Printf("Normal House WindowType: %v\n", scylladbDatabaseClient.Client.SDB)
	// fmt.Printf("Normal House DoorType: %v\n", normalHouse.doorType)
	// fmt.Printf("Normal House NumFloor: %v\n", normalHouse.floor)
}
