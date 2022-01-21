package model

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func Connect() (conn *pgx.Conn) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("connected to DB!")
	}
	return
}
