package server

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lukaszzieba/go-fun/internal/db"
)

func NewSqlcQueries(dbUrl string) *db.Queries {
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db := db.New(pool)

	return db
}
