package main

import "os"

func SetEnv() (err error) {

	var dbInfo = map[string]string{
		"DB_NAME": "gb_challenge",
		"DB_HOST": "localhost",
		"DB_USER": "postgres",
		"DB_PASS": "",
		"DB_PORT": "5432",
	}

	for key := range dbInfo {
		currEnv := os.Getenv(key)
		if currEnv != "" {
			dbInfo[key] = currEnv
		}
	}

	err = os.Setenv(
		"DB_URL",
		"host="+dbInfo["DB_HOST"]+
			" port="+dbInfo["DB_PORT"]+
			" user="+dbInfo["DB_USER"]+
			" dbname="+dbInfo["DB_NAME"]+
			" sslmode=disable",
	)
	return
}
