package model

import (
	"database/sql"
	"fmt"
)

type Question struct {
	Id           int    //`json:"id"`
	QuestionText string `db:"question_text"` //`json:"questionText"`
}

type Place struct {
	Country       string
	City          sql.NullString
	TelephoneCode int `db:"telcode"`
}

var queries = map[string]string{
	"insert":     "INSERT INTO questions (question_text) VALUES ($1) RETURNING *;",
	"getById":    "SELECT * FROM questions WHERE id=$1;",
	"getCurrent": "SELECT * FROM questions ORDER BY id DESC LIMIT 1;",
}

var id int
var questionText string

func QueryCurrentQuestion() (question Question, err error) {
	db := Open()
	err = db.QueryRowx(queries["getCurrent"]).StructScan(&question)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	return
}

/*
func QueryQuestionById(questionId int) (question Question, err error) {
	conn := Connect()
	err = conn.QueryRow(context.Background(), queries["getById"], questionId).Scan(&id, &questionText)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close(context.Background())
	return Question{id, questionText}, nil
}

func InsertQuestion(text string) (question Question, err error) {
	conn := Connect()
	err = conn.QueryRow(context.Background(), queries["insert"], text).Scan(&id, &questionText)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close(context.Background())
	return Question{id, questionText}, err
}

*/
