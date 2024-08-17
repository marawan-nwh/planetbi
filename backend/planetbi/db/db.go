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
	database := "planetbi"
	user := "marawan"
	pass := "marawan"
	host := "localhost"

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", user, pass, host, database)
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
