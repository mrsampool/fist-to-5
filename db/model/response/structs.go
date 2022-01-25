package response

import "encoding/json"

type QuestionResponses struct {
	Question  json.RawMessage `db:"question" json:"question"`
	Responses []Response      `json:"responses"`
}
type Response struct {
	Id      int `db:"id" json:"id"`
	Value   int `db:"value" json:"value"`
	Student Student
}
type Student struct {
	Id        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

type ResponseByQuestion struct {
	Id        int `db:"response_id" json:"id"`
	Value     int `db:"value" json:"value"`
	StudentId int `db:"student_id" json:"studentId"`
}

type NewResponse struct {
	QuestionId int `db:"question_id" json:"questionId"`
	StudentId  int `db:"student_id" json:"studentId"`
	Value      int `db:"value" json:"value"`
	ResponseId int `db:"response_id" json:"responseId"`
}
