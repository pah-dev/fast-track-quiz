/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		services.EndQuiz()
	},
}

func init() {
	rootCmd.AddCommand(endCmd)
}
