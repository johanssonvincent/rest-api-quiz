package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Answers struct {
	Answer1 string `json:"answer_1"`
	AnswerX string `json:"answer_x"`
	Answer2 string `json:"answer_2"`
}

type CorrectAnswers struct {
	Answer1 bool `json:"answer_1_correct"`
	AnswerX bool `json:"answer_x_correct"`
	Answer2 bool `json:"answer_2_correct"`
}

type Question struct {
	Question string `json:"question"`
	Answers Answers `json:"answers"`
	CorrectAnswers CorrectAnswers `json:"correct_answers"`
}

var questions = map[string]Question{
	"1": {
		Question: "What is the capital of France?",
		Answers: Answers{
			Answer1: "Paris",
			AnswerX: "London",
			Answer2: "Berlin",
		},
		CorrectAnswers: CorrectAnswers{
			Answer1: true,
			AnswerX: false,
			Answer2: false,
		},
	},
	"2": {
		Question: "What is the capital of Germany?",
		Answers: Answers{
			Answer1: "Paris",
			AnswerX: "London",
			Answer2: "Berlin",
		},
		CorrectAnswers: CorrectAnswers{
			Answer1: false,
			AnswerX: false,
			Answer2: true,
		},
	},
	"3": {
		Question: "What is the capital of England?",
		Answers: Answers{
			Answer1: "Paris",
			AnswerX: "London",
			Answer2: "Berlin",
		},
		CorrectAnswers: CorrectAnswers{
			Answer1: false,
			AnswerX: true,
			Answer2: false,
		},
	},
	"4": {
		Question: "What is the capital of Spain?",
		Answers: Answers{
			Answer1: "Paris",
			AnswerX: "London",
			Answer2: "Madrid",
		},
		CorrectAnswers: CorrectAnswers{
			Answer1: false,
			AnswerX: false,
			Answer2: true,
		},
	},
	"5": {
		Question: "What is the capital of Italy?",
		Answers: Answers{
			Answer1: "Paris",
			AnswerX: "Rome",
			Answer2: "Berlin",
		},
		CorrectAnswers: CorrectAnswers{
			Answer1: false,
			AnswerX: true,
			Answer2: false,
		},
	},
}

func main() {
	router := gin.Default()
	router.GET("/questions", getQuestions)
	router.GET("/questions/:id", getQuestion)

	router.Run("localhost:8080")
}

// getQuestions responds with the list of all questions as JSON.
func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

// getQuestion responds with the question for the specified ID.
func getQuestion(c *gin.Context) {
	id := c.Param("id")

	if question, ok := questions[id]; ok {
		c.IndentedJSON(http.StatusOK, question)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "question not found"})
	}
}
