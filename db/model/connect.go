package model

import (
	//"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"log"
	//"os"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func Open() (db *sqlx.DB) {
	/*
		db, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Println("connected to DB!")
		}
		return
	*/

	db, err := sqlx.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connected to DB!")
	}
	return
}
