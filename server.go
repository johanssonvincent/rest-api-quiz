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
	ID string `json:"id"`
	Question string `json:"question"`
	Answers Answers `json:"answers"`
	CorrectAnswers CorrectAnswers `json:"correct_answers"`
}

var questions = []Question{
	{
		ID: "1",
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
	{
		ID: "2",
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

	for _, question := range questions {
		if question.ID == id {
			c.IndentedJSON(http.StatusOK, question)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "question not found"})
}
