/*
Copyright © 2022 JEAN PAUL BOBENRIETH <jeanpaulb79 (at) gmail com>

*/
package cmd

import (
	"github.com/pah-dev/fast-track-quiz/cli/quiz/services"
	"github.com/spf13/cobra"
)

// endCmd represents the end command
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "End Game",
	Long: `
End Game

quiz end

End game to see your results.`,
	Run: func(cmd *cobra.Command, args []string) {
		services.EndQuiz()
	},
}

func init() {
	rootCmd.AddCommand(endCmd)
}
