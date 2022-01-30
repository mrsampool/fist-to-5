package db

import (
	//"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"log"
	//"os"
	_ "github.com/jackc/pgx/v4"
	// _ "github.com/lib/pq"
)

func Open() (db *sqlx.DB) {

	db, err := sqlx.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connected to DB!")
	}
	return
}
