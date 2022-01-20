package model

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func GetAllQuestions() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("connected to DB!")
	}
	defer conn.Close(context.Background())

	var id int
	var question string
	err = conn.QueryRow(context.Background(), "select * from questions where id=$1;", 1).Scan(&id, &question)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, question)
}
