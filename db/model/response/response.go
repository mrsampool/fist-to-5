package response

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"github.com/mrsampool/fist-to-5/db/model"
)

type QuestionResponses struct {
	Question  json.RawMessage `db:"question" json:"question"`
	Responses []Response      `json:"responses"`
}
type QuestionResponsesRaw struct {
	Question   json.RawMessage `db:"question"`
	ResList    []Response
	RawResList pq.StringArray `db:"responses"`
}
type Response struct {
	Id      int `db:"id" json:"id"`
	Value   int `db:"value" json:"value"`
	Student Student
}
type ResponseByQuestion struct {
	Id        int `db:"response_id" json:"id"`
	Value     int `db:"value" json:"value"`
	StudentId int `db:"student_id" json:"studentId"`
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

func QueryResponsesByQuestionId(questionId int) (responses []ResponseByQuestion, err error) {
	db := model.Open()
	rows, err := db.Queryx(Queries["queryByQuestion"], questionId)
	for rows.Next() {
		var response ResponseByQuestion
		err = rows.StructScan(&response)
		responses = append(responses, response)
	}
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	return
}

/*
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
