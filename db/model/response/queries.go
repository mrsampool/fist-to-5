package response

var Queries = map[string]string{
	"insert": `
		INSERT INTO responses 
		(student_id, question_id, value) 
		VALUES ($1, $2, $3) 
		RETURNING 
			question_id, 
			student_id, 
			value, 
			id AS response_Id;`,

	"queryCurrent": `
		WITH current_question AS (
			SELECT * 
			FROM questions 
			ORDER BY id DESC 
			LIMIT 1
		) 
		SELECT array(
			SELECT json_build_object(
				'student', (
					SELECT json_build_object(
						'id', student_id, 
						'first_name', first_name, 
						'last_name', last_name
					) FROM students 
					WHERE id=student_id
				), 
				'value', value, 
				'id', id
			) FROM responses 
			
		) AS responses, 
		(SELECT json_build_object(
				'id', id, 
				'text', question_text
			) AS question 
			FROM current_question
		);`,

	"queryByQuestion": `
		SELECT 
			student_id, 
			value, 
			id AS response_id 
		FROM responses 
		WHERE question_id=$1;`,
}
