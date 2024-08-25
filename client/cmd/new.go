/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"net/http"
	"bytes"
	"io"
	"fmt"

	"github.com/spf13/cobra"
)

type Question struct {
	QuestionAndAnswers QuestionAndAnswers `json:"question_and_answers"`
	CorrectAnswer string `json:"correct_answer"`
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new question to the quiz",
	Long: `Add a new question to the quiz`,
	Run: func(cmd *cobra.Command, args []string) {
		var q Question

		q.QuestionAndAnswers.Question, _ = cmd.Flags().GetString("question")
		q.QuestionAndAnswers.Answers.Answer1, _ = cmd.Flags().GetString("answer1")
		q.QuestionAndAnswers.Answers.AnswerX, _ = cmd.Flags().GetString("answerX")
		q.QuestionAndAnswers.Answers.Answer2, _ = cmd.Flags().GetString("answer2")
		
		// Chose to have user input 1, X, or 2 to make it simpler
		// Match the input to the correct answer
		correct, _ := cmd.Flags().GetString("correct")
		if correct == "1" {
			q.CorrectAnswer = q.QuestionAndAnswers.Answers.Answer1
		} else if correct == "2" {
			q.CorrectAnswer = q.QuestionAndAnswers.Answers.Answer2
		} else {
			q.CorrectAnswer = q.QuestionAndAnswers.Answers.AnswerX
		}

		url := "http://localhost:8080/questions"

		jsonData, err := json.Marshal(q)
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

		// Notify if creation failed
		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("Error: status code %d\n", resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Response body: %s\n", string(body))
			return
		}

		var decoded map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
			fmt.Printf("Failed to decode response: %v\n", err)
			return
		}

		fmt.Printf("Your question was successfully created with ID: %v\n", decoded["id"])
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("question", "q", "", "The the question")
	newCmd.Flags().StringP("answer1", "1", "", "Answer alternative 1")
	newCmd.Flags().StringP("answerX", "X", "", "Answer alternative X")
	newCmd.Flags().StringP("answer2", "2", "", "Answer alternative 2")
	newCmd.Flags().StringP("correct", "c", "", "The correct answer for the question (1, X, or 2)")
	
	newCmd.MarkFlagRequired("question")
	newCmd.MarkFlagRequired("answer1")
	newCmd.MarkFlagRequired("answerX")
	newCmd.MarkFlagRequired("answer2")
	newCmd.MarkFlagRequired("correct")

	newCmd.SetUsageTemplate(help)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Template for help message
var help = `Usage:		quiz new [flags]
Example:	quiz new -q "Question?" -1 "Yes" -X "Maybe" -2 "No" -c 1
	
Flags:
	-q, --question string		The question
	-1, --answer1 string		Answer alternative 1
	-X, --answerX string		Answer alternative X
	-2, --answer2 string		Answer alternative 2
	-c, --correct string		The correct answer for the question (1, X, or 2)
	-h, --help			help for new
`
