package server

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/lukaszzieba/az-go-test/internal/db"
)

func NewSqlcQueries(dbUrl string) *db.Queries {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to database")
	db := db.New(conn)

	return db
}
