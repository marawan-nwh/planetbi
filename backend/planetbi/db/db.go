package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Pool *pgxpool.Pool

var ErrNoRows = pgx.ErrNoRows

func Init() {
	host := "localhost"
	database := "planetbi"
	user := "marawan"

	// locally, posgresql setup sometimes saves the password somewhere and uses that for the connection,
	// so even if this is wrong, it will still work
	pass := "marawan"

	// dbURI := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", user, pass, host, database)
	dbURI := fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=%s", user, pass, host, database)
	config, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		panic(err)
	}

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	println("Connected to postgres 🔥")
}
