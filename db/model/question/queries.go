package question

var Queries = map[string]string{
	"insert":       "INSERT INTO questions (question_text) VALUES ($1) RETURNING *;",
	"getById":      "SELECT * FROM questions WHERE id=$1;",
	"queryCurrent": "SELECT * FROM questions ORDER BY id DESC LIMIT 1;",
	"queryCount":   "SELECT COUNT(*) FROM questions;",
	"deleteAll":    "DELETE * FROM questions;",
}
