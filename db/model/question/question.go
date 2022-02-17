package question

import (
	"fmt"
	"github.com/mrsampool/fist-to-5/db"
)

func QueryCurrentQuestion() (question Question, err error) {
	conn := db.Open()
	err = conn.QueryRowx(Queries["queryCurrent"]).StructScan(&question)
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

func QueryCount() (numQuestions int, err error) {
	type questionCount struct {
		Count int `db:"count"`
	}
	var count questionCount
	conn := db.Open()
	err = conn.QueryRowx(Queries["queryCount"]).StructScan(&count)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close()
	return count.Count, err
}

func DeleteAll() (err error) {
	conn := db.Open()
	defer conn.Close()
	_, err = conn.Exec(Queries["deleteAll"])
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	return
}
