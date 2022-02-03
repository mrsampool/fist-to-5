package question

import (
	"fmt"
	"github.com/mrsampool/fist-to-5/db"
)

func QueryCurrentQuestion() (question Question, err error) {
	conn := db.Open()
	err = conn.QueryRowx(Queries["getCurrent"]).StructScan(&question)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close()
	return
}

func QueryById(questionId int) (question Question, err error) {
	conn := db.Open()
	err = conn.QueryRowx(Queries["getById"], questionId).StructScan(&question)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close()
	return
}

func Insert(text string) (question Question, err error) {
	conn := db.Open()
	err = conn.QueryRowx(Queries["insert"], text).StructScan(&question)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close()
	return
}
