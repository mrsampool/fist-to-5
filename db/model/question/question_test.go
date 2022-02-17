package question

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	questionText := "Test question"
	newQuestion, err := Insert(questionText)
	if err != nil {
		t.Fatalf("Error inserting data: %d", err)
	}
	expectedType := "question.Question"
	actualType := reflect.TypeOf(newQuestion).String()
	if actualType != expectedType {
		t.Fatalf("Incorrect data format: expected %s but got %s", expectedType, actualType)
	}
}

func TestQueryCount(t *testing.T) {
	err := DeleteAll()
	if err != nil {
		t.Fatalf("Internal testing error: %s", err)
	}
	actualCount, err := QueryCount()
	actualType := reflect.TypeOf(actualCount).String()
	expectedType := "int"
	if actualType != expectedType {
		t.Fatalf("Incorrect data format: expected %s but got %s", expectedType, actualType)
	}
	expectedCount := 0
	if actualCount != expectedCount {
		t.Fatalf("Unexpected count: expected %s but got %s", expectedType, actualType)
	}
}
