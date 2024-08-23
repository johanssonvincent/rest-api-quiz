/*
Copyright © 2024 Vincent Johansson <vincent.johansson1@gmail.com>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// toplistCmd represents the toplist command
var toplistCmd = &cobra.Command{
	Use:   "toplist",
	Short: "Shows the top list of quiz results",
	Long: `Shows the top list of quiz results`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:8080/scores"
		var scores []Score
		getJson(url, &scores)

		printScores(scores)
	},
}

func init() {
	rootCmd.AddCommand(toplistCmd)
}

// Prints the banner for the toplist
func printBanner() {
	banner := `
╔════════════════════════════════════════╗
║                                        ║
║      ★ ★ ★ QUIZ LEADERBOARD ★ ★ ★      ║
║                                        ║
╚════════════════════════════════════════╝
`
	fmt.Println(banner)
}

// Prints the toplist of quiz results
func printScores(scores []Score) {
	printBanner()

	fmt.Println("╔" + strings.Repeat("═",40) + "╗")
	fmt.Printf("║ %-4s │ %-23s │ %-5s ║\n", "Rank", "Player", "Score")
	fmt.Println("╠" + strings.Repeat("═",40) + "╣")

	for i, score := range scores {
		if i == 10 {
			break
		}

		formattedName := formatName(score.Username)
		fmt.Printf("║ %-4d │ %-23s │ %3d   ║\n", i + 1, formattedName, score.Score)
	}
	fmt.Println("╚" + strings.Repeat("═",40) + "╝")
}

// Formats names longer than 23 characters to fit the high score table
func formatName(name string) string {
	if len(name) > 23 {
		return name[:20] + "..."
	} else {
		return name
	}
}