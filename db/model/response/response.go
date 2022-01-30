package response

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"github.com/mrsampool/fist-to-5/db"
)

func QueryResponsesByCurrentQuestion() (responses QuestionResponses, err error) {
	db := db.Open()
	var rawData struct {
		Question   json.RawMessage `db:"question"`
		ResList    []Response
		RawResList pq.StringArray `db:"responses"`
	}
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
	db := db.Open()
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

func InsertResponse(studentId, questionId, value int) (newResponse NewResponse, err error) {
	db := db.Open()
	err = db.Get(&newResponse, Queries["insert"], studentId, questionId, value)
	if err != nil {
		fmt.Println("QueryRow failed: ", err)
		return
	}
	fmt.Println(newResponse)
	return newResponse, err
}
