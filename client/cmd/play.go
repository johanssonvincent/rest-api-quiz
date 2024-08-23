/*
Copyright © 2024 Vincent Johansson <vincent.johansson1@gmail.com>

*/
package cmd

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"

	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
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

type userResult struct {
	Username string `json:"username"`
	Answers []string `json:"answers"`
}

type Score struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
	Percentage float64 `json:"percentage"`
}

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Starts the quiz",
	Long: `Starts the quiz`,

	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:8080/questions?type=short"
		var questions []QuestionAndAnswers
		getJson(url, &questions)

		answers := make([]string, len(questions))
		for i, q := range questions {
			answers[i] = answerQuestion(q)
		}

		url = "http://localhost:8080/scores"
		var result userResult

		// If the username flag is set, use that username
		if username, _ := cmd.Flags().GetString("username"); username != "" {
			result = userResult{Username: username, Answers: answers[:]}
		} else {
			prompt := promptui.Prompt{
				Label: "Enter your username to submit your answers",
			}
	
			username, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
	
			result = userResult{Username: username, Answers: answers[:]}
		}
		
		jsonData, err := json.Marshal(result)
		if err != nil {
			fmt.Printf("JSON marshalling failed %v\n", err)
			return
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("POST failed %v\n", err)
			return
		}

		defer resp.Body.Close()

		var score Score
		json.NewDecoder(resp.Body).Decode(&score)
		fmt.Printf("Good job, %s!\n", score.Username)
		if score.Percentage == 101 {
			fmt.Printf("Your score is %d out of %d and you were the first one to answer the quiz!\n", score.Score, len(questions))
		} else {
			fmt.Printf("Your score is %d out of %d, that's better than %.0f%% of quizzers!\n", score.Score, len(questions), score.Percentage)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	playCmd.Flags().StringP("username", "u", "", "Set your username to automatically submit your answers after finishing the quiz")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Displays a question with three possible answers and returns the selected answer
func answerQuestion(q QuestionAndAnswers) string {
	prompt := promptui.Select{
		Label: q.Question,
		Items: []string{q.Answers.Answer1, q.Answers.AnswerX, q.Answers.Answer2},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

// Reads JSON from a URL and decodes it into a target interface
func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}