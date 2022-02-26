/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/pah-dev/fast-track-quiz/cli/quiz/services"
	"github.com/spf13/cobra"
)

var qnCmd = &cobra.Command{
	Use:   "qn",
	Short: "Get a Question",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		services.GetOneQuestion()
	},
}

func init() {
	rootCmd.AddCommand(qnCmd)
}
