package response

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"github.com/mrsampool/fist-to-5/db/model"
)

type QuestionResponsesRaw struct {
	Question   json.RawMessage `db:"question"`
	ResList    []Response
	RawResList pq.StringArray `db:"responses"`
}
type QuestionResponses struct {
	Question  json.RawMessage `db:"question"`
	Responses []Response
}
type Response struct {
	Id      int `db:"id" json:"id"`
	Value   int `db:"id" json:"value"`
	Student Student
}
type Student struct {
	Id        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

func QueryResponsesByCurrentQuestion() (responses QuestionResponses, err error) {
	db := model.Open()
	var rawData QuestionResponsesRaw
	err = db.Get(&rawData, Queries["queryCurrent"])
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	responses.Question = rawData.Question
	for i, _ := range rawData.RawResList {
		var processedResponse Response
		err := json.Unmarshal([]byte(rawData.RawResList[i]), &processedResponse)
		if err != nil {
			fmt.Println(err)
		}
		responses.Responses = append(responses.Responses, processedResponse)
	}
	return
}

/*
func QueryResponsesByQuestionById(questionId int) (question Question, err error) {
	conn := Connect()
	err = conn.QueryRow(context.Background(), responseQueries["getById"], questionId).Scan(&id, &questionText)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close(context.Background())
	return Question{id, questionText}, nil
}

func InsertResponse(text string) (question Question, err error) {
	conn := Connect()
	err = conn.QueryRow(context.Background(), responseQueries["insert"], text).Scan(&id, &questionText)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close(context.Background())
	return Question{id, questionText}, err
}
*/