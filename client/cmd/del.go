/*
Copyright © 2024 Vincent Johansson <vincent.johansson1@gmail.com>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Deletes a question from the quiz",
	Long: `Deletes a question from the quiz`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:8080/questions/" + args[0]

		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			fmt.Printf("Error creating delete request: %v\n", err)
			return
		}
		
		// Send the delete request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending delete request: %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Failed to delete question with ID %s, please double check the ID\n", args[0])
			return
		}
		
		fmt.Printf("Question with ID %s deleted successfully\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	delCmd.SetUsageTemplate(`Usage:		quiz del [QUESTION_ID]
Example:	quiz del 3
	
Flags:
	-h, --help			help for 'del' command
`)
}
