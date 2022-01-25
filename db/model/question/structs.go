package question

type Question struct {
	Id           int    `db:"id" json:"id"`
	QuestionText string `db:"question_text" json:"text"`
}
