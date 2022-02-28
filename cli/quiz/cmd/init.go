/*
Copyright © 2022 JEAN PAUL BOBENRIETH <jeanpaulb79 (at) gmail com>

*/
package cmd

import (
	"github.com/pah-dev/fast-track-quiz/cli/quiz/services"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start Game",
	Long: `
Start the game
quiz init -n NAME
Write your name otherwise play as a guest.`,
	Run: func(cmd *cobra.Command, args []string) {
		if flag, _ := cmd.Flags().GetString("name"); flag != "" {
			services.StartQuiz(flag)
		}
	},
}

func init() {
	initCmd.Flags().StringP("name", "n", "Guest", "Player Name")
	rootCmd.AddCommand(initCmd)
}

