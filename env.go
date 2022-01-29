package main

import (
	"os"
)

func SetEnv() {
	os.Setenv("DB_URL", "host=localhost port=5432 user=sampool dbname=gb_challenge sslmode=disable")
}
