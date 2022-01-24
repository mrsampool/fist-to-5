package question

import (
	"database/sql"
	"fmt"
	"github.com/mrsampool/fist-to-5/db/model"
)

type Question struct {
	Id           int    `db:"id" json:"id"`
	QuestionText string `db:"question_text" json:"text"`
}

type Place struct {
	Country       string
	City          sql.NullString
	TelephoneCode int `db:"telcode"`
}

func QueryCurrentQuestion() (question Question, err error) {
	db := model.Open()
	err = db.QueryRowx(Queries["getCurrent"]).StructScan(&question)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	return
}

/*
func QueryQuestionById(questionId int) (question Question, err error) {
	conn := Connect()
	err = conn.QueryRow(context.Background(), Queries["getById"], questionId).Scan(&id, &questionText)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close(context.Background())
	return Question{id, questionText}, nil
}

func InsertQuestion(text string) (question Question, err error) {
	conn := Connect()
	err = conn.QueryRow(context.Background(), Queries["insert"], text).Scan(&id, &questionText)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close(context.Background())
	return Question{id, questionText}, err
}

*/
