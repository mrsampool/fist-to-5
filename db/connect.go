package db

import (
	"fmt"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
)

func Open() (db *sqlx.DB) {
	url, err := getUrl()
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err = sqlx.Open("postgres", url)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connected to DB!")
	}
	return
}

func getUrl() (url string, err error) {
	var dbInfo = map[string]string{
		"DB_NAME": "gb_challenge",
		"DB_HOST": "localhost",
		"DB_USER": "sampool",
		"DB_PASS": "",
		"DB_PORT": "5432",
	}
	if len(os.Args) > 1 && strings.Contains(os.Args[1], "test.v") {
		dbInfo["DB_NAME"] = "gb_challenge_test"
	}
	for key := range dbInfo {
		currEnv := os.Getenv(key)
		if currEnv != "" {
			dbInfo[key] = currEnv
		}
	}
	url = "host=" + dbInfo["DB_HOST"] +
		" port=" + dbInfo["DB_PORT"] +
		" user=" + dbInfo["DB_USER"] +
		" dbname=" + dbInfo["DB_NAME"] +
		" sslmode=disable"
	return
}
