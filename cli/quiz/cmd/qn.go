/*
Copyright © 2022 JEAN PAUL BOBENRIETH <jeanpaulb79 (at) gmail com>

*/
package cmd

import (
	"github.com/pah-dev/fast-track-quiz/cli/quiz/services"
	"github.com/spf13/cobra"
)

var qnCmd = &cobra.Command{
	Use:   "qn",
	Short: "Get a Question",
	Long: `
Get a Question

quiz qn

Get a question with four options.
You must answer the pending question, if you have one, before requesting another.`,
	Run: func(cmd *cobra.Command, args []string) {
		services.GetOneQuestion()
	},
}

func init() {
	rootCmd.AddCommand(qnCmd)
}
