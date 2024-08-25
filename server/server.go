package main

import (
	"net/http"
	"sort"
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
	QuestionAndAnswers QuestionAndAnswers `json:"question_and_answers"`
	CorrectAnswer string `json:"correct_answer"`
}

type UserResult struct {
	Username string `json:"username"`
	Answers map[int]string `json:"answers"`
}

type Score struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
	Percentage float64 `json:"percentage"`
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

var scores []Score

func main() {
	router := gin.Default()

	router.GET("/questions", getQuestions)
	router.GET("/questions/:id", getQuestion)
	router.POST("/questions", addQuestion)
	router.DELETE("/questions/:id", deleteQuestion)

	router.POST("/scores", postScore)
	router.GET("/scores", getScores)

	router.Run("localhost:8080")
}

// getQuestions responds with the list of all questions as JSON.
func getQuestions(c *gin.Context) {
	responseType := c.DefaultQuery("type", "default")

	switch responseType {
	case "short":
		questionList := map[int]QuestionAndAnswers{}

		for i, q := range questions {
			questionList[i] = q.QuestionAndAnswers
		}

		c.IndentedJSON(http.StatusOK, questionList)
	default:
		c.IndentedJSON(http.StatusOK, questions)
	}
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

// Adds a question to the map containing questions.
func addQuestion(c *gin.Context) {
	var question Question
	if err := c.BindJSON(&question); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid question"})
		return
	}

	if _, ok := questions[len(questions) * 2 + 1]; !ok {
		id := len(questions) * 2 + 1
		questions[id] = question
		c.IndentedJSON(http.StatusCreated, gin.H{
		"id": id,
		"message": "added successfully"})
		return
	}

	// Find lowest available ID and add question
	for i := 1; i < len(questions); i++ {
		if _, ok := questions[i]; !ok {
			questions[i] = question
			c.IndentedJSON(http.StatusCreated, gin.H{
				"id": i,
				"message": "added successfully"})
			return
		}
	}
}

// Deletes a question from the map containing questions.
func deleteQuestion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid ID"})
		return
	}

	if _, ok := questions[id]; ok {
		delete(questions, id)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "question deleted"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "question not found"})
	}
}

// postScore adds a score to the list of scores.
func postScore(c *gin.Context) {
	var userResult UserResult
	if err := c.BindJSON(&userResult); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid user result"})
		return
	}

	score := insertScore(userResult)

	if score != nil {
		c.IndentedJSON(http.StatusCreated, score)
		return
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid user result"})
	}
}

// getScores responds with the list of all scores as JSON.
func getScores(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, scores)
}

// insertScore adds a score to the list of scores in the correct position.
func insertScore(userResult UserResult) *Score {
	var newScore Score
	newScore.Username = userResult.Username
	newScore.Score = checkAnswers(userResult.Answers)

	index := sort.Search(len(scores), func(i int) bool {
		return scores[i].Score < newScore.Score
	})
	
	// Calculate the percentage of scores that are worse than the new score
	numScores := len(scores)
	worseCount := numScores - index	
	
	// Set percentage to 101 if it's the first score, avoids NaN value
	if numScores == 0 {
		newScore.Percentage = 101
	} else {
		newScore.Percentage = float64(worseCount) / float64(numScores) * 100
	}
	
	scores = append(scores[:index], append([]Score{newScore}, scores[index:]...)...)
	return &newScore
}

// Verifies users answers and return their score
func checkAnswers(answers map[int]string) int {
	score := 0

	for i, a := range answers {
		if a == questions[i].CorrectAnswer {
			score++
		}
	}

	return score
}