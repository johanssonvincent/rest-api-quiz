package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Answers struct {
	Answer1 string `json:"answer_1"`
	AnswerX string `json:"answer_x"`
	Answer2 string `json:"answer_2"`
}

type QuestionAndAnswers struct {
	Question string `json:"question"`
	Answers Answers `json:"answers"`
}

type Question struct {
	QuestionAndAnswers
	CorrectAnswer string `json:"correct_answer"`
}

var questions = map[int]Question{
	1: {
		QuestionAndAnswers: QuestionAndAnswers{
			Question: "What is the capital of France?",
			Answers: Answers{
				Answer1: "Paris",
				AnswerX: "London",
				Answer2: "Berlin",
			},
		},
		CorrectAnswer: "Paris",
	},
	2: {
		QuestionAndAnswers: QuestionAndAnswers{
			Question: "What is the capital of Germany?",
			Answers: Answers{
				Answer1: "Paris",
				AnswerX: "London",
				Answer2: "Berlin",
			},
		},
		CorrectAnswer: "Berlin",
	},
	3: {
		QuestionAndAnswers: QuestionAndAnswers{
			Question: "What is the capital of England?",
			Answers: Answers{
				Answer1: "Paris",
				AnswerX: "London",
				Answer2: "Berlin",
			},
		},
		CorrectAnswer: "London",
	},
	4: {
		QuestionAndAnswers: QuestionAndAnswers{
			Question: "What is the capital of Spain?",
			Answers: Answers{
				Answer1: "Paris",
				AnswerX: "London",
				Answer2: "Madrid",
			},
		},
		CorrectAnswer: "Madrid",
	},
	5: {
		QuestionAndAnswers: QuestionAndAnswers{
			Question: "What is the capital of Italy?",
			Answers: Answers{
				Answer1: "Paris",
				AnswerX: "Rome",
				Answer2: "Berlin",
			},
		},
		CorrectAnswer: "Rome",
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
	var questionList []QuestionAndAnswers

	for _, q := range questions {
		questionList = append(questionList, q.QuestionAndAnswers)
	}

	c.IndentedJSON(http.StatusOK, questionList)
}

// getQuestion responds with the question for the specified ID.
func getQuestion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid ID"})
		return
	}

	if question, ok := questions[id]; ok {
		c.IndentedJSON(http.StatusOK, question)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "question not found"})
	}
}
