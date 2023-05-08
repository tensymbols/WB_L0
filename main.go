package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func main() {
	//	const createTable = "CREATE TABLE users ( id serial not null, name VARCHAR(25) not null, birthday DATE not null);"
	//_, err = conn.Exec(context.Background(), sqlCreateTable)
	connString := "postgres://me:123@localhost:5432/wbdb"
	cfg, _ := pgx.ParseConfig(connString)
	ctx, _ := context.WithCancel(context.Background())
	logger := log.Default()

	conn, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		logger.Fatalf("unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

}
