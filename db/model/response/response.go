package response

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"github.com/mrsampool/fist-to-5/db"
)

func QueryResponsesByCurrentQuestion() (responses QuestionResponses, err error) {
	conn := db.Open()
	var rawData struct {
		Question   json.RawMessage `db:"question"`
		RawResList pq.StringArray  `db:"responses"`
	}
	err = conn.Get(&rawData, Queries["queryCurrent"])
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
	defer conn.Close()
	return
}

func QueryResponsesByQuestionId(questionId int) (responses []ResponseByQuestion, err error) {
	conn := db.Open()
	rows, err := conn.Queryx(Queries["queryByQuestion"], questionId)
	for rows.Next() {
		var response ResponseByQuestion
		err = rows.StructScan(&response)
		responses = append(responses, response)
	}
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close()
	return
}

func InsertResponse(studentId, questionId, value int) (newResponse NewResponse, err error) {
	conn := db.Open()
	err = conn.Get(&newResponse, Queries["insert"], studentId, questionId, value)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	defer conn.Close()
	return newResponse, err
}
