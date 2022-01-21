package model

import (
	"context"
	"fmt"
)

type Question struct {
	Id           int    `json:"id"`
	QuestionText string `json:"questionText"`
}

var queries = map[string]string{
	"insert":     "INSERT INTO questions (question_text) VALUES ($1) RETURNING *;",
	"getById":    "SELECT * FROM questions WHERE id=$1;",
	"getCurrent": "SELECT * FROM questions ORDER BY id DESC LIMIT 1;",
}

var id int
var questionText string

func QueryCurrentQuestion() (question Question, err error) {
	conn := Connect()
	err = conn.QueryRow(context.Background(), queries["getCurrent"]).Scan(&id, &questionText)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close(context.Background())
	return Question{id, questionText}, nil
}

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
