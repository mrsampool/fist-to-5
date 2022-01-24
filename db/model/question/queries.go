package question

var Queries = map[string]string{
	"insert":     "INSERT INTO questions (question_text) VALUES ($1) RETURNING *;",
	"getById":    "SELECT * FROM questions WHERE id=$1;",
	"getCurrent": "SELECT * FROM questions ORDER BY id DESC LIMIT 1;",
}
