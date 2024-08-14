/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/spf13/cobra"
	//"github.com/manifoldco/promptui"
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

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Starts the quiz",
	Long: `Starts the quiz`,

	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:8080/questions"
		var questions map[string]Question
		getJson(url, &questions)
		// Testing if json is parsed correctly 
		if value, exists := questions["1"]; exists {
			fmt.Printf("Question: %s\n", value.Question)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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