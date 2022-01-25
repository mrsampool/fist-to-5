package question

import (
	"fmt"
	"github.com/mrsampool/fist-to-5/db/model"
)

func QueryCurrentQuestion() (question Question, err error) {
	db := model.Open()
	err = db.QueryRowx(Queries["getCurrent"]).StructScan(&question)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	return
}

func QueryById(questionId int) (question Question, err error) {
	db := model.Open()
	err = db.QueryRowx(Queries["getById"], questionId).StructScan(&question)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	return
}

func Insert(text string) (question Question, err error) {
	db := model.Open()
	err = db.QueryRowx(Queries["insert"], text).StructScan(&question)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	return
}
